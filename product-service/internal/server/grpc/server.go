package grpc

import (
	"log"
	"net"
	"product-service/internal/config"
	"product-service/internal/repository"
	"product-service/internal/server/grpc/handler"

	v1 "github.com/Joshdike/ecom-microservices/common/proto/v1"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	config config.Config
}

func NewServer(config config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) StartGRPC() {

	conn := sqlx.MustConnect(
		"postgres",
		s.config.DBUrl,
	)

	// Create gRPC server
	grpcServer := grpc.NewServer()
	productPostGresRepos := repository.NewProductRepositoryPostgres(conn)
	productSvc := handler.NewProductGRPCHandler(productPostGresRepos)
	v1.RegisterProductServiceServer(grpcServer, productSvc)
	reflection.Register(grpcServer) // Enable gRPC reflection for testing

	// Start gRPC server

	grpcListener, err := net.Listen(s.config.Network, s.config.GrpcAddress)
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port: %v", err)
	}
	log.Println("starting grpc server on", s.config.GrpcAddress)
	if err := grpcServer.Serve(grpcListener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}

}
