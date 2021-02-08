package entity

type ProductDetail struct {
	Id                  string
	Description         string
	StockInformation    StockInformation
	CategoryInformation CategoryInformation
}

type StockInformation struct {
	AvailableCount int
}

type CategoryInformation struct {
	Id   int
	Name string
}
