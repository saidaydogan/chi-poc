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

type ProductModel struct {
	Name       string  `json:"name"`
	Sku        string  `json:"sku"`
	Price      float64 `json:"price"`
	CategoryId int     `json:"categoryId"`
}

type ProductDetailModel struct {
	Name     string         `json:"name"`
	Sku      string         `json:"sku"`
	Price    float64        `json:"price"`
	Category *CategoryModel `json:"category"`
}

type CategoryModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
