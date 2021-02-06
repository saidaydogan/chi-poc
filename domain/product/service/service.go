package service

import (
	"github.com/saidaydogan/chi-poc/domain/product/entity"
	"github.com/saidaydogan/chi-poc/domain/product/persistence"
)

type ProductService interface {
	GetProductById(id int) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id int) error
}

func NewProductService(repo persistence.ProductRepository) ProductService {

	return &productService{
		repo: repo,
	}

}

type productService struct {
	repo persistence.ProductRepository
}

func (r *productService) GetProductById(id int) (*entity.Product, error) {
	return r.repo.GetProductById(id)
}

func (r *productService) CreateProduct(product *entity.Product) error {
	return r.repo.CreateProduct(product)
}

func (r *productService) UpdateProduct(product *entity.Product) error {
	return r.repo.UpdateProduct(product)
}

func (r *productService) DeleteProduct(id int) error {
	return r.repo.DeleteProduct(id)
}
