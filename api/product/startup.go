package product

import (
	"github.com/go-chi/chi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/saidaydogan/chi-poc/api/product/controller"
	"github.com/saidaydogan/chi-poc/domain/product/service"
)

func Init(r chi.Router, productService service.ProductService, validator *validator.Validate, translator ut.Translator) {
	var handler = controller.NewBaseHandler(productService, validator, translator)

	r.Route("/products", func(r chi.Router) {
		r.Post("/", handler.Create)

		r.Route("/{productId}", func(r chi.Router) {
			r.Get("/", handler.GetById)
			r.Get("/", handler.GetById)
			r.Get("/detail", handler.GetDetailById)

			r.Put("/", handler.UpdateById)

			r.Delete("/", handler.DeleteById)
		})
	})
}
