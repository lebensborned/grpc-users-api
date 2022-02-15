package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	redis "github.com/go-redis/redis/v8"
	"github.com/lebensborned/grpc-crud/pkg/api"
)

type RedisClient struct {
	*redis.Client
}

func InitRedis(host, port string) *RedisClient {
	addrStr := fmt.Sprintf("%s:%s", host, port)
	cl := redis.NewClient(&redis.Options{
		Addr:     addrStr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	client := RedisClient{
		cl,
	}
	pong, err := cl.Ping(context.Background()).Result()
	log.Println("Redis ping: ", pong, err)
	return &client
}

func (c *RedisClient) SetUsers(data interface{}) error {
	marshal, err := json.Marshal(&data)
	if err != nil {
		log.Println("json marshal setUsers(): ", err)
		return err
	}
	err = c.Set(context.Background(), "users", string(marshal), -1).Err()
	log.Println("Redis-cache: users updated")
	if err != nil {
		log.Println("set users(): ", err)
		return err
	}
	return nil
}
func (c *RedisClient) GetUsers() ([]*api.UserProfile, error) {
	users := []*api.UserProfile{}
	bytes, err := c.Get(context.Background(), "users").Bytes()
	if err != nil {
		log.Println("getUsers() error: ", err)
		return nil, err
	}
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		log.Println("json unmarshal in getUsers(): ", err)
		return nil, err
	}
	return users, nil
}
