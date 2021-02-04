package controller

import (
	"github.com/go-chi/chi"
	"github.com/saidaydogan/chi-poc/domain/product/entity"
	"github.com/saidaydogan/chi-poc/domain/product/persistence"
	"net/http"
	"strconv"
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

	var (
		product *entity.Product
		err     error
	)

	idUrlParam := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idUrlParam)

	if product, err = c.productRepo.GetProductById(id); err != nil {
		errStatus := http.StatusInternalServerError

		if persistence.NotFoundError.Equal(err) {
			errStatus = http.StatusNotFound
		}

		respondWithError(w, errStatus, err.Error())
		return
	}

	respondwithJSON(w, http.StatusOK, product)
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
