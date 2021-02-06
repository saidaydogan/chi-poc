package persistence

import (
	"github.com/go-pg/pg"
	"github.com/saidaydogan/chi-poc/domain/product/entity"
)

type ProductRepository interface {
	GetProductById(id int) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id int) error
}

type productRepository struct {
	db *pg.DB
}

func NewProductRepository(db *pg.DB) ProductRepository {

	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetProductById(id int) (*entity.Product, error) {
	var product = entity.Product{
		Id: id,
	}

	err := r.db.Model(&product).WherePK().Select()
	if err == pg.ErrNoRows {
		return nil, NotFoundError
	}
	return &product, err
}

func (r *productRepository) CreateProduct(product *entity.Product) error {
	_, err := r.db.Model(product).Insert()
	return err
}

func (r *productRepository) UpdateProduct(product *entity.Product) error {
	_, err := r.db.Model(product).WherePK().Update()
	return err
}

func (r *productRepository) DeleteProduct(id int) error {
	var product = &entity.Product{
		Id: id,
	}
	_, err := r.db.Model(&product).Delete()
	return err
}
