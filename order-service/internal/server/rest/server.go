package rest

import (
	"log"
	"order-service/internal/config"
	"order-service/internal/repository"
	"order-service/internal/server/rest/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Server struct {
	config config.Config
}

func NewServer(config config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start() {

	r := gin.Default()
	// load it from env by calling the helper method
	cfg := config.Config{
		DbName: "somedb",
	}

	// put it into a helper function
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	repository.NewOrderRepositoryMongo(client, cfg)
	// Add order service
	orderHandler := handlers.NewOrderHandler()
	// create delete and get
	orderRoutes := r.Group("/orders", orderHandler.CreateOrder)
	// Create a db connection
	orderRoutes.POST("/")

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(s.config.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
