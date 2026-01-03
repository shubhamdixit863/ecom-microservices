package client

import (
	"context"
	v1 "github.com/Joshdike/ecom-microservices/common/proto/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductGrpcClient struct {
	conn          *grpc.ClientConn
	productclient v1.ProductServiceClient
}

// you can write a method invoke the server

func NewProductClient(address string) (*ProductGrpcClient, error) {
	conn, err := grpc.NewClient(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := v1.NewProductServiceClient(conn)

	return &ProductGrpcClient{
		conn:          conn,
		productclient: client,
	}, nil

}

func (c *ProductGrpcClient) GetProductById(ctx context.Context, productID int64) (*v1.ProductResponse, error) {
	req := &v1.ProductRequest{
		ProductId: productID,
	}

	return c.productclient.GetProductById(ctx, req)
}

func (c *ProductGrpcClient) Close() error {
	return c.conn.Close()
}
