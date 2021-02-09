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
	_ "github.com/saidaydogan/chi-poc/cmd/product-api/docs"
	"github.com/saidaydogan/chi-poc/domain/product/persistence"
	"github.com/saidaydogan/chi-poc/domain/product/service"
	"github.com/saidaydogan/chi-poc/pkg/db/postgre"
	"github.com/saidaydogan/chi-poc/pkg/httpswagger"
	"net/http"
	"time"
)

// @title Product API
// @version 1.0
// @description This is a sample REST API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://localhost:3333/
// @BasePath /v2
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

	r.Get("/swagger/*", httpswagger.WrapHandler)

	r.Route("/v1", func(r chi.Router) {
		productApi.Init(r, productService, validatorInstance, translator)
	})

	log.Info().Msg("Starting server...")

	err := http.ListenAndServe("localhost:3333", r)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
