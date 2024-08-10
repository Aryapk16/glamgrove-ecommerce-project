package main

import (
	"context"
	"glamgrove/glamgrove/productpb/glamgrove/productpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := productpb.NewProductServiceClient(conn)

	// Example: CreateProduct
	resp, err := c.CreateProduct(context.Background(), &productpb.CreateProductRequest{
		Name:        "New Product",
		Description: "A description of the new product",
		Price:       99.99,
	})
	if err != nil {
		log.Fatalf("could not create product: %v", err)
	}
	log.Printf("Created Product ID: %s", resp.Id)

	// Example: GetProduct
	prod, err := c.GetProduct(context.Background(), &productpb.GetProductRequest{
		Id: "some-product-id",
	})
	if err != nil {
		log.Fatalf("could not get product: %v", err)
	}
	log.Printf("Product: %s - %s - %f", prod.Name, prod.Description, prod.Price)
}
