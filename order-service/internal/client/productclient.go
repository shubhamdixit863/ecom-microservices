package client

import (
	productServiceGrpc "order-service/internal/client/proto/v1"

	"google.golang.org/grpc"
)

type ProductGrpcClient struct {
	conn          grpc.ClientConn
	productclient productServiceGrpc.ProductServiceClient
}

// you can write a method invoke the server

