package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	productApi "github.com/saidaydogan/chi-poc/api/product"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	productApi.Init(r)

	http.ListenAndServe(":3333", r)
}
