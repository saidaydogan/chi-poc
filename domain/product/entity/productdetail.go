package entity

type ProductDetail struct {
	Description         string
	StockInformation    StockInformation
	CategoryInformation CategoryInformation
}

type StockInformation struct {
	AvailableCount int
}

type CategoryInformation struct {
	Id   string
	Name string
}
