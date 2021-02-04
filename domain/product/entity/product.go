package entity

type Product struct {
	Id         int     `sql:"Id,pk"`
	Name       string  `sql:"Name"`
	Sku        string  `sql:"Sku"`
	Price      float64 `sql:"Price"`
	CategoryId int     `sql:"CategoryId"`
}
