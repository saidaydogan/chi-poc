package entity

type Category struct {
	Id          int    `pg:"Id,pk"`
	Name        string `pg:"Name"`
	Description string `pg:"Description"`
}
