package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	productApi "github.com/saidaydogan/chi-poc/api/product"
	"github.com/saidaydogan/chi-poc/domain/product/persistence"
	"github.com/saidaydogan/chi-poc/domain/product/service"
	"github.com/saidaydogan/chi-poc/pkg/db/postgre"
	"net/http"
	"time"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	var db = postgre.Initialize("postgres", "changeme", "product_db")

	productRepo := persistence.NewProductRepository(db)
	productService := service.NewProductService(productRepo)

	validatorInstance := validator.New()

	tEn := en.New()
	uni := ut.New(tEn, tEn)
	translator, _ := uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(validatorInstance, translator)

	r.Route("/v1", func(r chi.Router) {
		productApi.Init(r, productService, validatorInstance, translator)
	})

	log.Info().Msg("Starting server...")

	err := http.ListenAndServe("localhost:3333", r)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
