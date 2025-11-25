package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("product with id: %d", id),
	})
}

func (handler *ProductHandler) CreateProduct(c *gin.Context) {
	//A struct will hold the new product
	//bind the info in request to the struct, save to the database and return response

	c.JSON(http.StatusAccepted, gin.H{
		"message": "product created successfully",
	})
}

func (handler *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("product %d updated", id),
	})
}

func (handler *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("product %d deleted", id),
	})

}
