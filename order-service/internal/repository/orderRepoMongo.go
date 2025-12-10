package repository

import (
	"context"
	"order-service/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepositoryMongo struct {
	collection *mongo.Collection
}

func NewOrderRepositoryMongo(db *mongo.Database) OrderRepository {
	collection := db.Collection("orders")
	return &OrderRepositoryMongo{
		collection,
	}
}

func (o *OrderRepositoryMongo) CreateOrder(order models.Order) (string, error) {
	order.CreatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := o.collection.InsertOne(ctx, order)
	if err != nil {
		return "", err
	}
	objectID := result.InsertedID.(primitive.ObjectID)
	return objectID.Hex(), nil
}

func (o *OrderRepositoryMongo) GetOrder(id string) (models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Order{}, err
	}
	var order models.Order

	err = o.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (o *OrderRepositoryMongo) DeleteOrder(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = o.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}
