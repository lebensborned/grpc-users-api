package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	"github.com/jinzhu/gorm"
	"github.com/lebensborned/grpc-crud/config"
	"github.com/lebensborned/grpc-crud/pkg/api"
	"github.com/lebensborned/grpc-crud/pkg/cache"
	"github.com/lebensborned/grpc-crud/pkg/clickhouse"
	"github.com/lebensborned/grpc-crud/pkg/kafka"
	"github.com/lebensborned/grpc-crud/pkg/user"
	"github.com/lebensborned/grpc-crud/storage"
	"google.golang.org/grpc"
)

func main() {

	log.Println("Server starting...")

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	var (
		ch *clickhouse.ClickHouse
		db *gorm.DB
	)
	// waiting for clickhouse availability
	for {
		ch, err = clickhouse.Connect("clickhouse", config.CHname, config.CHpass, config.CHuser)
		if err == nil {
			break
		}
		log.Println("clickhouse reconnecting after 10 sec, error occured: ", err)
		time.Sleep(10 * time.Second)
	}
	// waiting for postgres availability
	for {
		db, err = storage.ConnectDB(config.DBhost, config.DBuser, config.DBname, config.DBpass, config.DBport)
		if err == nil {
			break
		}
		log.Println("postgres reconnecting after 10 sec, error occured: ", err)
		time.Sleep(10 * time.Second)
	}
	defer db.Close()

	client := cache.InitRedis(config.RedisHost, config.RedisPort)
	defer client.Close()

	// set up grpc api server
	s := grpc.NewServer()
	srv := &user.UserServer{DB: db, Cache: client, Config: config}
	api.RegisterUserProfilesServer(s, srv)
	l, err := net.Listen("tcp", config.ServerPort)
	if err != nil {
		log.Fatal("cannot listen tcp: ", err)
	}
	log.Println("Listening TCP on ", config.ServerPort)

	// cache users from postgres every min
	go func() {
		for {
			userlist := []*api.UserProfile{}
			err = srv.DB.Model(&api.UserProfile{}).Limit(100).Find(&userlist).Error
			if err != nil {
				log.Println("cant fetch users from db: ", err)
			}
			err = srv.Cache.SetUsers(userlist)
			if err != nil {
				log.Println("cant setusers in goroutine: ", err)
			}
			time.Sleep(1 * time.Minute)
		}

	}()

	// set up kafka consumer
	worker, err := kafka.ConnectConsumer([]string{"kafka:9092"})
	if err != nil {
		log.Fatal("cannot start consuming: ", err)
	}
	consumer, err := worker.ConsumePartition("users", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("cannot consume partition: ", err)
	}
	log.Println("Consumer started")
	defer consumer.Close()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	doneCh := make(chan struct{})

	// consumer loop
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println(err)
			case msg := <-consumer.Messages():
				log.Printf("Received message: Topic(%s) | Message(%s) \n", string(msg.Topic), string(msg.Value))
				ch.InsertData(string(msg.Value))
			case <-sigchan:
				doneCh <- struct{}{}
			}
		}
	}()
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
	<-doneCh
	if err := worker.Close(); err != nil {
		log.Fatal(err)
	}
}
