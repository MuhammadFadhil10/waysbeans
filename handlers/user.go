package handlers

import (
	"net/http"
	"waysbeans/helper"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/golang-jwt/jwt/v4"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	var user models.User
	var err error

	user, err = h.UserRepository.GetProfile(user, userId)
	if err != nil {
		helper.ResponseHelper(w, err, nil, http.StatusInternalServerError)
		return
	}

	userResponse := models.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Photo:    user.Photo,
		Role:     user.Role,
	}

	helper.ResponseHelper(w, nil, userResponse, 0)

}
