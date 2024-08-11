package main

import (
	"context"
	"glamgrove/glamgrove/productpb"

	"log"
	"net"

	"google.golang.org/grpc"
)

type productServer struct {
	productpb.UnimplementedProductServiceServer
}

// Implement the product service methods
func (s *productServer) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	return &productpb.CreateProductResponse{Id: "new-product-id"}, nil
}

func (s *productServer) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	return &productpb.GetProductResponse{Name: "Example Product", Price: 99.99}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server is listening on port 50052")

	s := grpc.NewServer()
	productpb.RegisterProductServiceServer(s, &productServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
