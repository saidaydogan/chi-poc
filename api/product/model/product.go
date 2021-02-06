package model

type CreateProductRequest struct {
	Name       string  `json:"name" validate:"required"`
	Sku        string  `json:"sku" validate:"required"`
	Price      float64 `json:"price" validate:"required,gte=1"`
	CategoryId int     `json:"categoryId" validate:"required,gte=1"`
}

type UpdateProductRequest struct {
	Name       string  `json:"name" validate:"required"`
	Sku        string  `json:"sku" validate:"required"`
	Price      float64 `json:"price" validate:"required,gte=1"`
	CategoryId int     `json:"categoryId" validate:"required,gte=1"`
}
