package handler

import (
	"context"
	"log"
	"product-service/internal/repository"
	v1 "github.com/Joshdike/ecom-microservices/common/proto/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductGRPCHandler struct {
	productRepository repository.ProductRepository
	v1.UnimplementedProductServiceServer
}

func NewProductGRPCHandler(productRepository repository.ProductRepository) *ProductGRPCHandler {
	return &ProductGRPCHandler{productRepository: productRepository}
}

func (s *ProductGRPCHandler) GetProductById(ctx context.Context, req *v1.ProductRequest) (*v1.ProductResponse, error) {
	log.Println("Product Grpc Handler Get Product By Id invoked")

	product, err := s.productRepository.GetProductByID(req.ProductId)
	if err != nil {
		log.Printf("Error getting product: %v", err)
		return nil, status.Error(codes.NotFound, "product not found")
	}

	return &v1.ProductResponse{
		ProductId:   product.ID,
		ProductName: product.Name,
	}, nil
}
