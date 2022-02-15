package main

import (
	"context"
	"log"

	"github.com/lebensborned/grpc-crud/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Client started")
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewUserProfilesClient(conn)
	u := &api.UserProfile{
		Name: "John",
		Age:  321,
	}
	res, err := c.CreateUserProfile(context.Background(), &api.CreateUserProfileRequest{UserProfile: u})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", res)
	rez, err := c.ListUserProfiles(context.Background(), &api.EmptyReq{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", rez.Profiles)
}
