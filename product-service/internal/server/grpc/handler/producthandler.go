package handler

import (
	"context"
	"product-service/internal/server/grpc"
)

type ProductHandler struct {
	grpc.ProductServiceServer
}

func (p *ProductHandler) GetProductById(context.Context, *grpc.ProductRequest) (*grpc.ProductResponse, error) {
	// We will write logic here
	// here we will make the db call as well to get the data
}
