package product

import (
	"github.com/go-chi/chi"
	"github.com/saidaydogan/chi-poc/api/product/controller"
	"github.com/saidaydogan/chi-poc/domain/product/persistence"
)

func Init(r chi.Router, productRepo persistence.ProductRepository) {
	var handler = controller.NewBaseHandler(productRepo)

	r.Route("/products", func(r chi.Router) {
		r.Route("/{productId}", func(r chi.Router) {
			r.Get("/", handler.GetById)
			r.Get("/detail", handler.GetDetailById)

			r.Put("/", handler.UpdateById)

			r.Delete("/", handler.DeleteById)
		})
	})
}
