package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	dto "waysbeans/dto/result"
	"waysbeans/helper"
	jwttoken "waysbeans/pkg/jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Status: "Failed", Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwttoken.DecodeToken(token)

		if err != nil {
			helper.ResponseHelper(w, err, nil, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userInfo", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
