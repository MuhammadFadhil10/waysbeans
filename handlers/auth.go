package handlers

import (
	"encoding/json"
	"net/http"
	dto "waysbeans/dto/result"
	"waysbeans/helper"
	"waysbeans/models"
	bcryptpkg "waysbeans/pkg/bcrypt"
	"waysbeans/repositories"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	decodeErr := json.NewDecoder(r.Body).Decode(&user)
	helper.ResponseHelper(w, decodeErr, nil, http.StatusBadRequest, false)

	userExist, userErr := h.AuthRepository.GetByEmail(user, user.Email)

	if userErr == nil {
		w.WriteHeader(http.StatusConflict)
		response := dto.ErrorResult{Status: "error", Message: "Email " + userExist.Email + "already registered"}
		json.NewEncoder(w).Encode(response)
		return
	}

	hashedPassword, hashedPasswordErr := bcryptpkg.HashingPassword(user.Password)
	helper.ResponseHelper(w, hashedPasswordErr, nil, http.StatusInternalServerError, false)
	user.Password = hashedPassword

	var err error
	user, err = h.AuthRepository.CreateUser(user)

	helper.ResponseHelper(w, err, user, http.StatusInternalServerError, true)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userRequest models.User

	decodeErr := json.NewDecoder(r.Body).Decode(&userRequest)
	helper.ResponseHelper(w, decodeErr, nil, http.StatusBadRequest, false)

	userLogin, err := h.AuthRepository.LoginUser(userRequest, userRequest.Email)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		response := dto.ErrorResult{Status: "error", Message: "Email not registered!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	passwordMatch := bcryptpkg.CheckPasswordHash(userRequest.Password, userLogin.Password)

	if !passwordMatch {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "error", Message: "Password wrong!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	helper.ResponseHelper(w, nil, userLogin, 0, true)

}
