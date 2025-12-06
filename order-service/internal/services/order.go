package services

import "order-service/internal/repository"

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService() repository.OrderRepository {
	return &OrderService{}
}
