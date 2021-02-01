package persistence

import "github.com/saidaydogan/chi-poc/domain/product/entity"

type ProductRepository interface {
	GetProductById(id int) entity.Product
	CreateProduct(product entity.Product)
	UpdateProduct(id int, product entity.Product)
	DeleteProduct(id int)
}

type productRepository struct {
}

func (r *productRepository) GetProductById(id int) entity.Product {
	return entity.Product{}
}

func (r *productRepository) CreateProduct(product entity.Product) {

}

func (r *productRepository) UpdateProduct(id int, product entity.Product) {

}

func (r *productRepository) DeleteProduct(id int) {
}

func newProductRepository() *productRepository {

	return &productRepository{}
}
