package main

import (
	"glamgrove/glamgrove/productpb/glamgrove/productpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type productServer struct {
	productpb.UnimplementedProductServiceServer
}

// Implement your server methods here

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
