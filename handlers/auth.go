package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	dto "waysbeans/dto/result"
	"waysbeans/helper"
	"waysbeans/models"
	bcryptpkg "waysbeans/pkg/bcrypt"
	jwttoken "waysbeans/pkg/jwt"
	"waysbeans/repositories"

	"github.com/golang-jwt/jwt/v4"
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
	if decodeErr != nil {
		helper.ResponseHelper(w, decodeErr, nil, http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.FullName == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "error", Message: "Fill all register field!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	userExist, userErr := h.AuthRepository.GetByEmail(user, user.Email)

	if userErr == nil {
		w.WriteHeader(http.StatusConflict)
		response := dto.ErrorResult{Status: "error", Message: "Email " + userExist.Email + "already registered"}
		json.NewEncoder(w).Encode(response)
		return
	}

	hashedPassword, hashedPasswordErr := bcryptpkg.HashingPassword(user.Password)
	if hashedPasswordErr != nil {
		helper.ResponseHelper(w, hashedPasswordErr, nil, http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	userPhoto := os.Getenv("PATH_FILE") + "default_profile.png"

	user.Photo = userPhoto

	var err error
	user, err = h.AuthRepository.CreateUser(user)

	helper.ResponseHelper(w, err, user, http.StatusInternalServerError)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userRequest models.User

	decodeErr := json.NewDecoder(r.Body).Decode(&userRequest)
	if decodeErr != nil {
		helper.ResponseHelper(w, decodeErr, nil, http.StatusBadRequest)
		return
	}

	var user models.User
	var err error
	user, err = h.AuthRepository.GetByEmail(user, userRequest.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: "error", Message: "You are not registered!"}
		json.NewEncoder(w).Encode(response)
		return
	}

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

	generateToken := jwt.MapClaims{}

	generateToken["id"] = userLogin.ID
	generateToken["exp"] = time.Now().Add(time.Hour * 3).Unix()

	token, _ := jwttoken.CreateToken(&generateToken)

	resp := map[string]models.UserLoginResponse{
		"user": {
			Email: userLogin.Email,
			Photo: userLogin.Photo,
			Token: token,
			Role:  userLogin.Role,
		},
	}

	helper.ResponseHelper(w, nil, resp, 0)

}

func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	var user models.User
	user, _ = h.AuthRepository.GetAuthProfile(user, userId)

	generateToken := jwt.MapClaims{}

	// generateToken["id"] = userLogin.ID
	generateToken["exp"] = time.Now().Add(time.Hour * 3).Unix()

	token, _ := jwttoken.CreateToken(&generateToken)

	resp := models.UserLoginResponse{
		Email: user.Email,
		Photo: user.Photo,
		Token: token,
	}
	helper.ResponseHelper(w, nil, resp, 0)
}
