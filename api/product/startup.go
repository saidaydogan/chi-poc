package product

import (
	"github.com/go-chi/chi"
	"github.com/saidaydogan/chi-poc/api/product/controller"
)

func Init(r *chi.Mux) {
	r.Route("/product", func(r chi.Router) {
		r.Get("/{id}", controller.GetById)
		r.Get("/{id}/detail", controller.GetDetailById)
	})
}
