package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (handler *ProductHandler) GetProducts(c *gin.Context) {
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "all products here",
	})
}

func (handler *ProductHandler) GetProductByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "product one",
	})
}

func (handler *ProductHandler) CreateProduct(c *gin.Context){

}

func(handler *ProductHandler) UpdateProduct(c *gin.Context){

}

func(handler *ProductHandler) DeleteProduct(c *gin.Context){
	
}
