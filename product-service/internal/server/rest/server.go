package rest

import (
	"log"
	"product-service/internal/repository"
	"product-service/internal/server/rest/handlers"
	"product-service/internal/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	productRepo repository.ProductRepository
}

func NewServer(productRepo repository.ProductRepository) *Server {
	return &Server{
		productRepo,
	}
}

func (s *Server) Start() {

	productSvc := services.NewProductService(s.productRepo)
	productHandler := handlers.NewProductHandler(productSvc)

	// Create a Gin router with default middleware (logger and recovery)

	r := gin.Default()

	productRoutes := r.Group("/products")
	// Create a db connection

	// Define a simple GET endpoint
	productRoutes.GET("/get-product", productHandler.GetProducts)
	productRoutes.POST("/create", productHandler.CreateProduct)
	productRoutes.GET("/get-product/:id", productHandler.GetProductByID)
	productRoutes.PUT("/update/:id", productHandler.UpdateProduct)
	productRoutes.DELETE("/delete/:id", productHandler.DeleteProduct)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":8092"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
