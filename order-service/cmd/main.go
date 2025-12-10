package main

import (
	"log"
	"order-service/internal/config"
	"order-service/internal/server/rest"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Load .env from parent directory
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file from parent directory")
	}
	dbUrl := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("MONGO_DB")

	cnf := config.NewConfig(":8090", dbUrl, dbName)
	s := rest.NewServer(cnf)
	s.Start()

}
