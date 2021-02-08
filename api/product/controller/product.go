package controller

import (
	"encoding/json"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/saidaydogan/chi-poc/api/product/model"
	"github.com/saidaydogan/chi-poc/domain/product/entity"
	"github.com/saidaydogan/chi-poc/domain/product/persistence"
	"github.com/saidaydogan/chi-poc/domain/product/service"
	"github.com/saidaydogan/chi-poc/pkg/validatorhelper"
	"net/http"
)

type BaseHandler struct {
	productService service.ProductService
	validator      *validator.Validate
	translator     ut.Translator
}

func NewBaseHandler(productService service.ProductService, validator *validator.Validate, translator ut.Translator) *BaseHandler {
	return &BaseHandler{
		productService: productService,
		validator:      validator,
		translator:     translator,
	}
}

func (c *BaseHandler) Create(w http.ResponseWriter, r *http.Request) {

	var (
		createRequest model.CreateProductRequest
		product       *entity.Product
	)

	if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.validator.Struct(createRequest); err != nil {
		respondWithErrors(w, http.StatusUnprocessableEntity, validatorhelper.ToErrResponse(err, c.translator).Errors)
		return
	}

	//product = &entity.Product{
	//	Name:       createRequest.Name,
	//	Sku:        createRequest.Sku,
	//	Price:      createRequest.Price,
	//	CategoryId: createRequest.CategoryId,
	//}

	if err := c.productService.CreateProduct(product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondwithJSON(w, http.StatusCreated, product)
}

func (c *BaseHandler) GetById(w http.ResponseWriter, r *http.Request) {

	var (
		productEntity *entity.Product
		productModel  *model.ProductModel
		id            int
	)

	id = getUrlParamInt(r, "productId")
	productEntity = getProductById(c, w, id)
	if productEntity == nil {
		return
	}

	productModel = &model.ProductModel{
		Name:       productEntity.Name,
		Sku:        productEntity.Sku,
		Price:      productEntity.Price,
		CategoryId: productEntity.CategoryId,
	}

	respondwithJSON(w, http.StatusOK, productModel)
}

func (c *BaseHandler) GetDetailById(w http.ResponseWriter, r *http.Request) {
	var (
		productEntity      *entity.Product
		productDetailModel *model.ProductDetailModel
		id                 int
		err                error
	)

	id = getUrlParamInt(r, "productId")
	if productEntity, err = c.productService.GetProductDetailById(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	productDetailModel = &model.ProductDetailModel{
		Name:  productEntity.Name,
		Sku:   productEntity.Sku,
		Price: productEntity.Price,
		Category: &model.CategoryModel{
			Id:   productEntity.Category.Id,
			Name: productEntity.Category.Name,
		},
	}

	respondwithJSON(w, http.StatusOK, productDetailModel)
}

func (c *BaseHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	var (
		updateRequest model.UpdateProductRequest
		product       *entity.Product
		id            int
		err           error
	)

	id = getUrlParamInt(r, "productId")
	product = getProductById(c, w, id)
	if product == nil {
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.validator.Struct(updateRequest); err != nil {
		respondWithErrors(w, http.StatusUnprocessableEntity, validatorhelper.ToErrResponse(err, c.translator).Errors)
		return
	}

	product.Name = updateRequest.Name
	product.Sku = updateRequest.Sku
	product.Price = updateRequest.Price
	//product.CategoryId = updateRequest.CategoryId

	if err = c.productService.UpdateProduct(product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondwithJSON(w, http.StatusOK, product)

}

func (c *BaseHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteById"))
}

func getProductById(c *BaseHandler, w http.ResponseWriter, id int) *entity.Product {
	var (
		product *entity.Product
		err     error
	)

	if product, err = c.productService.GetProductById(id); err != nil {
		errStatus := http.StatusInternalServerError

		if persistence.NotFoundError.Equal(err) {
			errStatus = http.StatusNotFound
		}

		respondWithError(w, errStatus, err.Error())
		return nil
	}
	return product
}
