package repository

import (
	"context"
	"order-service/internal/config"
	"order-service/internal/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type OrderRepositoryMongo struct {
	client *mongo.Client
	config config.Config
}

func NewOrderRepositoryMongo(client *mongo.Client, config config.Config) OrderRepository {
	return OrderRepositoryMongo{
		client: client,
		config: config,
	}
}

// Make sure to implement the interface
// load the data from config in env
func (o *OrderRepositoryMongo) CreateOrder(order models.Order) error {
	_, err := o.client.Database(o.config.DbName).Collection("order").
		InsertOne(context.Background(), bson.D{{"product_name", order.ProductName}, {"quantity", order.ProductQuantity}})
	if err != nil {
		return err
	}

	return nil
}
