package model

type ProductDetail struct {
	Description         string              `json:"description"`
	StockInformation    StockInformation    `json:"stockInformation"`
	CategoryInformation CategoryInformation `json:"categoryInformation"`
}

type StockInformation struct {
	AvailableCount int `json:"availableCount"`
}

type CategoryInformation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
