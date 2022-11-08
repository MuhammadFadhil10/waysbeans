package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	productdto "waysbeans/dto/products"
	dto "waysbeans/dto/result"
	"waysbeans/helper"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

type handler struct {
	ProductRepository repositories.ProductsRepository
}

func HandlerProduct(ProductRepository repositories.ProductsRepository) *handler {
	return &handler{ProductRepository}
}

func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var products []models.Products
	var err error

	products, err = h.ProductRepository.GetProducts(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if len(products) < 1 {
		w.WriteHeader(http.StatusNotFound)

		errorMessage := "record not found"
		response := dto.ErrorResult{Status: "error", Message: errorMessage}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: convertProductResponse(products)}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) CreateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stock, _ := strconv.Atoi(r.FormValue("stock"))
	price, _ := strconv.Atoi(r.FormValue("price"))

	request := productdto.CreateProductRequest{
		Name:        r.FormValue("name"),
		Stock:       stock,
		Price:       price,
		Description: r.FormValue("description"),
		Photo:       r.FormValue("photo"),
	}

	product := models.Products{
		Name:        request.Name,
		Stock:       request.Stock,
		Price:       request.Price,
		Description: request.Description,
		Photo:       request.Photo,
	}

	product, err := h.ProductRepository.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: product}
	json.NewEncoder(w).Encode(response)

}

func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	productId, _ := strconv.Atoi(mux.Vars(r)["productId"])

	var product models.Products

	product, err := h.ProductRepository.GetProductById(product, productId)

	helper.ResponseHelper(w, err, product)

}

func convertProductResponse(p []models.Products) map[string][]models.Products {
	var products []models.Products

	for _, product := range p {
		products = append(products, product)
	}

	resp := map[string][]models.Products{
		"products": products,
	}

	return resp
}
