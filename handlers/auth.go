package handlers

import (
	"encoding/json"
	"net/http"
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

	hashedPassword, hashedPasswordErr := bcryptpkg.HashingPassword(user.Password)
	helper.ResponseHelper(w, hashedPasswordErr, nil, http.StatusInternalServerError, false)
	user.Password = hashedPassword

	var err error
	user, err = h.AuthRepository.CreateUser(user)

	helper.ResponseHelper(w, err, user, http.StatusInternalServerError, true)

}
