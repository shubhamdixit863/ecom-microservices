package repository

import "order-service/internal/models"

type OrderRepository interface {
	CreateOrder(order models.Order) (string, error)
	GetOrder(id string) (models.Order, error)
	GetOrders() ([]models.Order, error)
	DeleteOrder(id string) error
}
