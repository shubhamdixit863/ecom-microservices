package repository

import "product-service/internal/models"

type ProductRepository interface {
	CreateProduct(product models.Product) (int64, error)
	GetProductByID(id int64) (models.Product, error)
	GetProducts()([]models.Product, error)
	UpdateProduct(id int64, update models.Product) error
	DeleteProduct(id int64) error
}
