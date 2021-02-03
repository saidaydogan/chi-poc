package product

import (
	"github.com/go-chi/chi"
	"github.com/go-pg/pg"
	"github.com/saidaydogan/chi-poc/api/product/controller"
	"github.com/saidaydogan/chi-poc/domain/product/persistence"
)

func Init(r chi.Router) {
	var db *pg.DB
	productRepo := persistence.NewProductRepository(db)

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
