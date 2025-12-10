package grpc

import (
	"log"
	"net"
	"product-service/internal/config"
	"product-service/internal/repository"
	v1 "product-service/internal/server/grpc/v1"
	"product-service/internal/services"

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
	productSvc := services.NewProductGRPCService(productPostGresRepos)
	v1.RegisterProductServiceServer(grpcServer, productSvc)
	reflection.Register(grpcServer) // Enable gRPC reflection for testing

	// Start gRPC server
	grpcListener, err := net.Listen(s.config.Network, s.config.GrpcAddress)
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port: %v", err)
	}
	log.Println("Grpc server has been started")
	if err := grpcServer.Serve(grpcListener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}

}
