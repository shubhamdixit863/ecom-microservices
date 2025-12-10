package handlers

import (
	"net/http"
	"order-service/internal/dto"
	"order-service/internal/services"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (p *OrderHandler) CreateOrder(c *gin.Context) {
	var orderRequest dto.OrderRequest
	err := c.BindJSON(&orderRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	id, err := p.orderService.CreateOrder(orderRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "order created sucessfully",
		"id":      id,
	})
}

func (p *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := p.orderService.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func (p *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	err := p.orderService.DeleteOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "order deleted successfully",
	})
}
