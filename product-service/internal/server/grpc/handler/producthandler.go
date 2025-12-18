package handler

import (
	"context"
	"product-service/internal/repository"
	"product-service/internal/server/grpc"
)

type ProductHandler struct {
	grpc.ProductServiceServer
	productRepo repository.ProductRepository
}

func NewProductHandler(productRepo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{productRepo: productRepo}
}

func (p *ProductHandler) GetProductById(ctx context.Context, productRequest *grpc.ProductRequest) (*grpc.ProductResponse, error) {
	// We will write logic here
	// here we will make the db call as well to get the data
	// we will do database call and get the data

	product, err := p.productRepo.GetProductByID(productRequest.ProductId)
	if err != nil {
		return nil, err
	}

	return &grpc.ProductResponse{
		ProductId:   product.ID,
		ProductName: product.Name,
	}, nil
}
