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

	var id int
	query := "INSERT INTO products (name, price, description) VALUES ($1, $2, $3) RETURNING id"
	err := p.db.QueryRow(query, product.Name, product.Price, product.Description).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
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
