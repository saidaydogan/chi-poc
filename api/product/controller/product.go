package controller

import (
	"fmt"
	"github.com/saidaydogan/chi-poc/domain/product/entity"
	"github.com/saidaydogan/chi-poc/domain/product/persistence"
	"net/http"
)

type BaseHandler struct {
	productRepo persistence.ProductRepository
}

func NewBaseHandler(productRepo persistence.ProductRepository) *BaseHandler {
	return &BaseHandler{
		productRepo: productRepo,
	}
}

func (c *BaseHandler) GetById(w http.ResponseWriter, r *http.Request) {

	var product *entity.Product

	if product, err := c.productRepo.GetProductById(1); err != nil {
		fmt.Println("Error", product)
	}

	w.Write([]byte(fmt.Sprintf("GetById %s", product.Name)))
}

func (c *BaseHandler) GetDetailById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetDetailById"))
}

func (c *BaseHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateById"))
}

func (c *BaseHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteById"))
}
