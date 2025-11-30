package handlers

import (
	"fmt"
	"net/http"
	"product-service/internal/dto"
	"product-service/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (handler *ProductHandler) GetProducts(c *gin.Context) {
	products, err := handler.productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"count": len(products),
	})
}

func (handler *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	product, err := handler.productService.GetProductByID(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

// dto  --->thats basically yor request response
// model  --->which you will save inside db

func (handler *ProductHandler) CreateProduct(c *gin.Context) {

	var productRequest dto.ProductRequest

	err := c.BindJSON(&productRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return

	}

	//A struct will hold the new product
	//bind the info in request to the struct, save to the database and return response
	id, err := handler.productService.CreateProduct(productRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "product created successfully",
		"id":      id,
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
