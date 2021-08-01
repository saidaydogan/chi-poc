package persistence

import (
	"errors"
	"github.com/saidaydogan/chi-poc/domain/product/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductById(id int) (*entity.Product, error)
	GetProductDetailById(id int) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {

	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetProductById(id int) (*entity.Product, error) {
	var product = &entity.Product{
		Id: id,
	}

	result := r.db.First(product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, NotFoundError
	}
	return product, result.Error
}

func (r *productRepository) GetProductDetailById(id int) (*entity.Product, error) {
	var product = &entity.Product{
		Id: id,
	}

	result := r.db.Debug().Joins("Category").First(product)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, NotFoundError
	}
	return product, result.Error
}

func (r *productRepository) CreateProduct(product *entity.Product) error {
	result := r.db.Create(product)
	return result.Error
}

func (r *productRepository) UpdateProduct(product *entity.Product) error {
	result := r.db.Save(&product)
	return result.Error
}

func (r *productRepository) DeleteProduct(id int) error {
	var product = &entity.Product{
		Id: id,
	}
	result := r.db.Delete(product)
	return result.Error
}
