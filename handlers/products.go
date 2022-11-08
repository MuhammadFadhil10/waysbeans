package handlers

import (
	"encoding/json"
	"net/http"
	dto "waysbeans/dto/result"
	"waysbeans/models"
	"waysbeans/repositories"
)

type handler struct {
	ProductRepository repositories.ProductsRepository
}

func HandlerProduct(ProductRepository repositories.ProductsRepository) *handler {
	return &handler{ProductRepository}
}

func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var productsModels []models.Products

	products, err := h.ProductRepository.GetProducts(productsModels)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if len(products) <= 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	response := dto.SuccessResult{Status: "success", Data: convertProductResponse(products)}
	json.NewEncoder(w).Encode(response)
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
