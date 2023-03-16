package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/skyfoxs/api.simple-twitter/pkg/handler"
)

type ContextKey string

const UserID ContextKey = "user"

type JWTAuth struct {
	Secret []byte
}

func (m JWTAuth) TokenRequired(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("Authorization")
		s := strings.Split(a, "Bearer ")
		if len(s) != 2 {
			handler.Unauthorized(w, r)
			return
		}
		t, err := jwt.Parse(s[1], func(token *jwt.Token) (interface{}, error) {
			return m.Secret, nil
		})
		if err != nil {
			handler.Unauthorized(w, r)
			return
		}

		if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
			ctx := context.WithValue(r.Context(), UserID, claims["userId"])
			f.ServeHTTP(w, r.WithContext(ctx))
		} else {
			handler.Unauthorized(w, r)
		}
	}
}
