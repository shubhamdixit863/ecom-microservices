package services

import (
	"order-service/internal/dto"
	"order-service/internal/models"
	"order-service/internal/repository"
)

type OrderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{orderRepo: repo}
}

func (s *OrderService) CreateOrder(orderRequest dto.OrderRequest) (string, error) {
	order := models.Order{
		ProductName:        orderRequest.ProductName,
		ProductDescription: orderRequest.ProductDescription,
		ProductQuantity:    orderRequest.ProductQuantity,
		Address:            orderRequest.Address,
	}
	id, err := s.orderRepo.CreateOrder(order)
	if err != nil{
		return "", err
	}
	return id, nil
}

func (s *OrderService) GetOrder(id string) (dto.OrderResponse, error) {
	order, err := s.orderRepo.GetOrder(id)
	if err != nil{
		return dto.OrderResponse{}, err
	}
	OrderResponse := dto.OrderResponse{
		ID: order.ID,
		ProductName: order.ProductName,
		ProductDescription: order.ProductDescription,
		ProductQuantity: order.ProductQuantity,
		Address: order.Address,
	}
	return OrderResponse, nil
}

func (s *OrderService) DeleteOrder(id string) error {
	err := s.orderRepo.DeleteOrder(id)
	if err != nil{
		return err
	}
	return nil
}
