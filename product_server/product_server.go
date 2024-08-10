// product_server.go
package main

import (
	"context"
	"log"
	"net"

	"GlamGrove/proto/glamgrove/productpb" // Adjust import path based on directory structure

	"google.golang.org/grpc"
)

type productServer struct {
	productpb.UnimplementedProductServiceServer
}

func (s *productServer) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	// Implement your logic here
	return &productpb.CreateProductResponse{Id: "new-product-id"}, nil
}

func (s *productServer) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	// Implement your logic here
	return &productpb.GetProductResponse{Name: "Example Product", Description: "Product Description", Price: 9.99}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	productpb.RegisterProductServiceServer(s, &productServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
