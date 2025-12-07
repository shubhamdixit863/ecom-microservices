package main

import (
	"log"
	"os"
	"product-service/internal/repository"
	"product-service/internal/server/grpc/handler"
	"product-service/internal/server/rest"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env from parent directory
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file from parent directory")
	}

	dbUrl := os.Getenv("DATABASE_URL")

	//// Instantiate the server object
	//cnf := config.Config{
	//	Port:  ":8090",
	//	DBUrl: dbUrl,
	//}

	conn := sqlx.MustConnect(
		"postgres",
		dbUrl,
	) // Dependency injection
	productPostGresRepos := repository.NewProductRepositoryPostgres(conn)
	// We will start with the grpc server
	go func() {
		handler.StartGRPC(productPostGresRepos)

	}()

	s := rest.NewServer(productPostGresRepos)
	s.Start()

}
