package client

import (
	productServiceGrpc "order-service/internal/client/proto"

	"google.golang.org/grpc"
)

type ProductGrpcClient struct {
	conn          grpc.ClientConn
	productclient productServiceGrpc.ProductServiceClient
}

// you can write amethiod invoke the server
