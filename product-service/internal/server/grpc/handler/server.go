package handler

import (
	"log"
	"net"
	"product-service/internal/repository"
	grpc2 "product-service/internal/server/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPC(productRepo repository.ProductRepository) {

	// Create gRPC server
	grpcServer := grpc.NewServer()
	producttGrpcHandler := NewProductHandler(productRepo)
	grpc2.RegisterProductServiceServer(grpcServer, producttGrpcHandler)
	reflection.Register(grpcServer) // Enable gRPC reflection for testing

	// Start gRPC server
	grpcListener, err := net.Listen("tcp", "localhost:50090")
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port: %v", err)
	}
	log.Println("Grpc server has been started")
	if err := grpcServer.Serve(grpcListener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}

}
