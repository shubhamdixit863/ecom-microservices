package services

import (
	"context"
	"log"
	"product-service/internal/repository"
	v1 "product-service/internal/server/grpc/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductGRPCService struct {
	productRepository repository.ProductRepository
	v1.UnimplementedProductServiceServer
}

func NewProductGRPCService(productRepository repository.ProductRepository) *ProductGRPCService {
	return &ProductGRPCService{productRepository: productRepository}
}

func (s *ProductGRPCService) GetProductById(ctx context.Context, req *v1.ProductRequest) (*v1.ProductResponse, error) {
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
