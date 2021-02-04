package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	productApi "github.com/saidaydogan/chi-poc/api/product"
	"github.com/saidaydogan/chi-poc/domain/product/persistence"
	"github.com/saidaydogan/chi-poc/pkg/db/postgre"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	var db = postgre.Initialize("postgres", "changeme", "product_db")
	productRepo := persistence.NewProductRepository(db)

	r.Route("/v1", func(r chi.Router) {
		productApi.Init(r, productRepo)
	})

	http.ListenAndServe(":3333", r)
}
