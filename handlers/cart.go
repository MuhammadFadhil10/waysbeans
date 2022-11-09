package handlers

import (
	"encoding/json"
	"net/http"
	dto "waysbeans/dto/result"
	"waysbeans/helper"
	"waysbeans/models"
	"waysbeans/repositories"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cart := models.Cart{}

	err := json.NewDecoder(r.Body).Decode(&cart)
	helper.ResponseHelper(w, err, nil, http.StatusInternalServerError, false)

	cart, err = h.CartRepository.AddToCart(cart)
	helper.ResponseHelper(w, err, cart, http.StatusBadRequest, true)

}

func (h *handlerCart) GetCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")

	var carts []models.Cart

	carts, err := h.CartRepository.GetCarts(carts)
	if err == nil {
		if len(carts) == 0 {
			w.WriteHeader(http.StatusNotFound)
			errorMessage := "Carts not found!"
			response := dto.ErrorResult{Status: "error", Message: errorMessage}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	helper.ResponseHelper(w, err, carts, http.StatusInternalServerError, true)
}
