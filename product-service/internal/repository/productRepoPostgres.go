package repository

import (
	"product-service/internal/models"

	"github.com/jmoiron/sqlx"
)

type ProductRepositoryPostgres struct {
	db *sqlx.DB
}

func (p *ProductRepositoryPostgres) CreateProduct(product models.Product) (int, error) {
	// We will write the sql query
	_, err := p.db.DB.Exec("INSERT INTO products(name,price,description) VALUES ($1,$2,$3) RETURNING id", product.Name, product.Price, product.Description)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (p *ProductRepositoryPostgres) GetProductByID() {
	//TODO implement me
	panic("implement me")
}

func (p *ProductRepositoryPostgres) GetProducts() {
	//TODO implement me
	panic("implement me")
}

func NewProductRepositoryPostgres(conn *sqlx.DB) ProductRepository {
	return &ProductRepositoryPostgres{
		conn,
	}
}
