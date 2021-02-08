package persistence

import (
	"github.com/go-pg/pg/v10"
	"github.com/saidaydogan/chi-poc/domain/product/entity"
)

type ProductRepository interface {
	GetProductById(id int) (*entity.Product, error)
	GetProductDetailById(id int) (*entity.Product, error)
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
	var product = &entity.Product{
		Id: id,
	}
	//var product = new(entity.Product)

	err := r.db.Model(product).WherePK().Select()
	if err == pg.ErrNoRows {
		return nil, NotFoundError
	}
	return product, err
}

func (r *productRepository) GetProductDetailById(id int) (*entity.Product, error) {
	var product = entity.Product{
		Id: id,
	}

	// Join columns
	//err := r.db.Model(&product).Relation("Category", func(q *orm.Query) (*orm.Query, error) {
	//	join := q.TableModel().GetJoin("Category")
	//	join.Columns = []string{"Id", "Name"}
	//	return q, nil
	//}).WherePK().Select()
	//

	err := r.db.Model(&product).Relation("Category").WherePK().Select()
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
