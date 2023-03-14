package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const ContextUserIDKey ContextKey = "user"

func respondUnauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(ErrorResponse{Message: "Unauthorized access"})
}

func requiredToken(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("Authorization")
		s := strings.Split(a, "Bearer ")
		if len(s) != 2 {
			respondUnauthorized(w)
			return
		}
		t, err := jwt.Parse(s[1], func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil {
			respondUnauthorized(w)
			return
		}

		if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
			ctx := context.WithValue(r.Context(), ContextUserIDKey, claims["userId"])
			f.ServeHTTP(w, r.WithContext(ctx))
		} else {
			respondUnauthorized(w)
		}
	}
}
