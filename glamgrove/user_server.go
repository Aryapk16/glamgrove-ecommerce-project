package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := userpb.NewUserServiceClient(conn)

	// Example: CreateUser
	resp, err := c.CreateUser(context.Background(), &userpb.CreateUserRequest{
		Name:  "New User",
		Email: "newuser@example.com",
	})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("Created User ID: %s", resp.Id)

	// Example: GetUser
	usr, err := c.GetUser(context.Background(), &userpb.GetUserRequest{
		Id: "some-user-id",
	})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User: %s - %s", usr.Name, usr.Email)
}
