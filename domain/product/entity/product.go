package entity

import "github.com/saidaydogan/chi-poc/domain/category/entity"

type Product struct {
	tableName  struct{}         `pg:"products,discard_unknown_columns"`
	Id         int              `pg:"Id, pk"`
	Name       string           `pg:"Name"`
	Sku        string           `pg:"Sku"`
	Price      float64          `pg:"Price"`
	CategoryId int              `pg:"CategoryId"`
	Category   *entity.Category `pg:"rel:has-one, fk:CategoryId"`
}
