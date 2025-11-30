package services

import (
	"product-service/internal/dto"
	"product-service/internal/models"
	"product-service/internal/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) CreateProduct(productRequest dto.ProductRequest) (int64, error) {

	// We will convert the dto to model
	product := models.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
	}
	id, err := s.productRepository.CreateProduct(product)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func(s *ProductService) GetProducts() ([]dto.ProductResponse, error){
	products, err := s.productRepository.GetProducts()
	if err != nil{
		return nil, err
	}
	var productRes []dto.ProductResponse
	for _, product := range products{
		res := dto.ProductResponse{
			ID: product.ID,
			Name: product.Name,
			Price: product.Price,
			Description: product.Description,
		}
		productRes = append(productRes, res)
	}
	return productRes, nil
}

func (s *ProductService) GetProductByID(id int64)(dto.ProductResponse, error){
	product, err := s.productRepository.GetProductByID(id)
	if err != nil{
		return dto.ProductResponse{}, err
	}
	productRes := dto.ProductResponse{
		ID: product.ID,
		Name: product.Name,
		Price: product.Price,
		Description: product.Description,
	}
	return productRes, nil
}

func (s *ProductService) UpdateProduct(id int64, update dto.ProductRequest)error{
	product := models.Product{
		Name: update.Name,
		Price: update.Price,
		Description: update.Description,
	}

	err := s.productRepository.UpdateProduct(id, product)
	if err != nil{
		return err
	}
	return nil
}

func (s *ProductService) DeleteProduct(id int64)error{
	err := s.productRepository.DeleteProduct(id)
	if err != nil{
		return err
	}
	return nil
}