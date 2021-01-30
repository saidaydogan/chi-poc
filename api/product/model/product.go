package model

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Sku   string  `json:"sku"`
	Price float64 `json:"price"`
}
