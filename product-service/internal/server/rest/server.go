package rest

import (
	"log"
	"product-service/internal/config"
	"product-service/internal/repository"
	"product-service/internal/server/rest/handlers"
	"product-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	conn := sqlx.MustConnect(
		"postgres",
		s.config.DBUrl,
	) // Dependency injection
	productPostGresRepos := repository.NewProductRepositoryPostgres(conn)
	productSvc := services.NewProductService(productPostGresRepos)
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
	if err := r.Run(s.config.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
