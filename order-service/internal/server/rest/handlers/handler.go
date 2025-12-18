package handlers

import "github.com/gin-gonic/gin"

type OrderHandler struct {
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (p *OrderHandler) CreateOrder(c *gin.Context) {

}
