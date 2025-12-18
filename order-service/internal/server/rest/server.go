package rest

import (
	"context"
	"log"
	productGrpcClient "order-service/internal/client"
	"order-service/internal/config"
	"order-service/internal/repository"
	"order-service/internal/server/rest/handlers"
	"order-service/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start() {

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.config.DBUrl))
	if err != nil {
		log.Fatalf("failed to connect to mongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB")

	// Get database
	db := client.Database(s.config.DbName)

	orderRepo := repository.NewOrderRepositoryMongo(db)
	productClient, err := productGrpcClient.NewProductClient("localhost:50090")
	if err != nil {
		log.Fatal("Error", err)
		return
	}
	ordersvc := services.NewOrderService(orderRepo, productClient)
	orderHandler := handlers.NewOrderHandler(ordersvc)

	r := gin.Default()
	orderRoutes := r.Group("/orders")

	orderRoutes.POST("/create", orderHandler.CreateOrder)
	orderRoutes.GET("/get-order", orderHandler.GetOrders)
	orderRoutes.GET("/get-order/:id", orderHandler.GetOrder)
	orderRoutes.DELETE("/delete/:id", orderHandler.DeleteOrder)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(s.config.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
