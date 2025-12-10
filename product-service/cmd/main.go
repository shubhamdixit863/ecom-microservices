package main

import (
	"log"
	"os"
	"product-service/internal/config"
	"product-service/internal/server/grpc"
	"product-service/internal/server/rest"

	"github.com/joho/godotenv"
)

func main() {

	// Load .env from parent directory
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file from parent directory")
	}

	dbUrl := os.Getenv("DATABASE_URL")

	// Instantiate the server object
	cnf := config.Config{
		Port:        ":8090",
		DBUrl:       dbUrl,
		Network:     "tcp",
		GrpcAddress: "localhost:50090",
	}

	go func() {
		grpc.NewServer(cnf).StartGRPC()
	}()

	s := rest.NewServer(cnf)
	s.Start()

}
