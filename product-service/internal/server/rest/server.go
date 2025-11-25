package rest

import (
	"log"
	"product-service/internal/config"
	"product-service/internal/server/rest/handlers"

	"github.com/gin-gonic/gin"
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

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	productRoutes := r.Group("/products")
	productHandler := handlers.NewProductHandler()

	// Define a simple GET endpoint
	productRoutes.GET("/get-product", productHandler.GetProducts)
	productRoutes.POST("/create", productHandler.CreateProduct)
	productRoutes.GET("/get-product/:id", productHandler.GetProductByID)
	productRoutes.PUT("/update/:id", productHandler.UpdateProduct)
	productRoutes.DELETE("/delete/:id", productHandler.DeleteProduct)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(s.config.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
