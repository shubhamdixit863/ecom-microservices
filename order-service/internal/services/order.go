package services

import (
	"context"
	"log"
	"order-service/internal/client"
	"order-service/internal/dto"
	"order-service/internal/models"
	"order-service/internal/repository"
)

type OrderService struct {
	orderRepo                repository.OrderRepository
	productServiceGrpcClient *client.ProductGrpcClient
}

func NewOrderService(repo repository.OrderRepository, productServiceGrpcClient *client.ProductGrpcClient) *OrderService {
	return &OrderService{orderRepo: repo, productServiceGrpcClient: productServiceGrpcClient}
}

func (s *OrderService) CreateOrder(orderRequest dto.OrderRequest) (string, error) {
	order := models.Order{
		ProductName:        orderRequest.ProductName,
		ProductDescription: orderRequest.ProductDescription,
		ProductQuantity:    orderRequest.ProductQuantity,
		Address:            orderRequest.Address,
	}
	// We will invoke the product service to check whether the order is correct or not
	// whether we have proper quantity of the product or not

	byId, err := s.productServiceGrpcClient.GetProductById(context.Background(), orderRequest.ProductID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Println("The product", byId)

	id, err := s.orderRepo.CreateOrder(order)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *OrderService) GetOrder(id string) (dto.OrderResponse, error) {
	order, err := s.orderRepo.GetOrder(id)
	if err != nil {
		return dto.OrderResponse{}, err
	}
	OrderResponse := dto.OrderResponse{
		ID:                 order.ID,
		ProductName:        order.ProductName,
		ProductDescription: order.ProductDescription,
		ProductQuantity:    order.ProductQuantity,
		Address:            order.Address,
	}
	return OrderResponse, nil
}

func (s *OrderService) GetOrders() ([]dto.OrderResponse, error) {
	var orderResponse []dto.OrderResponse
	orders, err := s.orderRepo.GetOrders()
	if err != nil {
		return orderResponse, err
	}

	for _, order := range orders {
		OrderResponse := dto.OrderResponse{
			ID:                 order.ID,
			ProductName:        order.ProductName,
			ProductDescription: order.ProductDescription,
			ProductQuantity:    order.ProductQuantity,
			Address:            order.Address,
		}
		orderResponse = append(orderResponse, OrderResponse)
	}

	return orderResponse, nil
}

func (s *OrderService) DeleteOrder(id string) error {
	err := s.orderRepo.DeleteOrder(id)
	if err != nil {
		return err
	}
	return nil
}
