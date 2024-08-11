package main

import (
	"context"
	"fmt"
	"glamgrove/glamgrove/productpb"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := productpb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateProduct(ctx, &productpb.CreateProductRequest{Name: "Test Product"})
	if err != nil {
		log.Fatalf("could not create product: %v", err)
	}
	fmt.Printf("Product ID: %s\n", r.GetId())
}
