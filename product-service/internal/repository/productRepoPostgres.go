package repository

import (
	"fmt"
	"product-service/internal/models"

	"github.com/jmoiron/sqlx"
)

type ProductRepositoryPostgres struct {
	db *sqlx.DB
}

func (p *ProductRepositoryPostgres) CreateProduct(product models.Product) (int64, error) {
	// We will write the sql query

	var id int64
	query := "INSERT INTO products (name, price, description) VALUES ($1, $2, $3) RETURNING id"
	err := p.db.QueryRow(query, product.Name, product.Price, product.Description).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *ProductRepositoryPostgres) GetProductByID(id int64) (models.Product, error) {
	query := "SELECT * FROM products WHERE id = $1"
	var product models.Product
	err := p.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price,
		&product.Description, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (p *ProductRepositoryPostgres) GetProducts() ([]models.Product, error) {
	query := "SELECT * FROM products"
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting products: %w", err)
	}

	defer rows.Close()
	var products []models.Product

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description,
			&product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductRepositoryPostgres) UpdateProduct(id int64, update models.Product) error {
	query := `UPDATE products 
    SET 
        name = COALESCE(NULLIF($1, ''), name),
        price = CASE WHEN $2 != 0.00 THEN $2 ELSE price END,
        description = COALESCE(NULLIF($3, ''), description)
    WHERE id = $4`

	_, err := p.db.Exec(query, update.Name, update.Price, update.Description, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepositoryPostgres) DeleteProduct(id int64) error {
	query := `DELETE FROM products WHERE id = $1`

	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepositoryPostgres(conn *sqlx.DB) ProductRepository {
	return &ProductRepositoryPostgres{
		conn,
	}
}
