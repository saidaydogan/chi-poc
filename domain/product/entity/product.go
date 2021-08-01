package entity

import "github.com/saidaydogan/chi-poc/domain/category/entity"

type Product struct {
	Id         int
	Name       string
	Sku        string
	Price      float64
	CategoryId int
	Category   *entity.Category
}
