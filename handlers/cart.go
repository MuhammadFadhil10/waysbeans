package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	dto "waysbeans/dto/result"
	"waysbeans/helper"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	cart := models.Cart{}

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		helper.ResponseHelper(w, err, nil, http.StatusInternalServerError)
		return
	}

	cartExist, err := h.CartRepository.GetCartExist(userId, cart.ProductID)

	if err == nil {
		cartExist.TotalPrice = cartExist.TotalPrice + (cartExist.TotalPrice / cartExist.Qty)
		cartExist.Qty = cartExist.Qty + 1
		cart, err = h.CartRepository.UpdateCartQty(cartExist, cartExist.ID)
	} else {
		cart.UserID = userId
		cart, err = h.CartRepository.AddToCart(cart)
	}

	if err != nil {
		helper.ResponseHelper(w, err, cart, http.StatusBadRequest)
		return
	}
	helper.ResponseHelper(w, err, cart, http.StatusBadRequest)

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
	helper.ResponseHelper(w, err, carts, http.StatusInternalServerError)
}

func (h *handlerCart) GetCartByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	var carts []models.Cart
	var err error

	carts, err = h.CartRepository.GetCartByUser(carts, userId)
	if err != nil {
		helper.ResponseHelper(w, err, nil, http.StatusInternalServerError)
		return
	}

	helper.ResponseHelper(w, err, carts, 0)

}

func (h *handlerCart) UpdateCartQty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	updateType := r.URL.Query()["update"]
	cartId, _ := strconv.Atoi(mux.Vars(r)["cartId"])

	var cartUpdate models.Cart
	var err error
	cartUpdate, err = h.CartRepository.GetCart(cartUpdate, cartId)

	if len(updateType) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "error", Message: "Missing query parameter (update)"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if updateType[0] == "add" {
		cartUpdate.TotalPrice = cartUpdate.TotalPrice + (cartUpdate.TotalPrice / cartUpdate.Qty)
		cartUpdate.Qty = cartUpdate.Qty + 1

	} else {
		cartUpdate.TotalPrice = cartUpdate.TotalPrice - (cartUpdate.TotalPrice / cartUpdate.Qty)
		cartUpdate.Qty = cartUpdate.Qty - 1
	}
	cartUpdate.UpdateAt = time.Now()

	cart, err := h.CartRepository.UpdateCartQty(cartUpdate, cartId)

	helper.ResponseHelper(w, err, cart, http.StatusInternalServerError)

}

func (h *handlerCart) DeleteCartByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cartId, _ := strconv.Atoi(mux.Vars(r)["cartId"])

	var cartDeleted models.Cart
	var err error
	cartDeleted, err = h.CartRepository.DeleteCartByID(cartDeleted, cartId)

	response := map[string]models.Cart{
		"cartDeleted": cartDeleted,
	}

	if err != nil {
		helper.ResponseHelper(w, err, response, http.StatusInternalServerError)
		return
	}

	helper.ResponseHelper(w, err, response, http.StatusInternalServerError)

}

func (h *handlerCart) DeleteCartByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])

	var cartDeleted models.Cart

	err := h.CartRepository.DeleteCartByUser(cartDeleted, userId)

	if err != nil {
		helper.ResponseHelper(w, err, cartDeleted, http.StatusInternalServerError)
		return
	}

	helper.ResponseHelper(w, err, cartDeleted, http.StatusInternalServerError)
}
