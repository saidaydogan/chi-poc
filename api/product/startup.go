package product

import (
	"github.com/go-chi/chi"
	"github.com/saidaydogan/chi-poc/api/product/controller"
)

func Init(r chi.Router) {
	r.Route("/products", func(r chi.Router) {
		r.Route("/{productId}", func(r chi.Router) {
			r.Get("/", controller.GetById)
			r.Get("/detail", controller.GetDetailById)

			r.Put("/", controller.UpdateById)

			r.Delete("/", controller.DeleteById)
		})
	})
}
