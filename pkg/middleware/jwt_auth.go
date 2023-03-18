package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/skyfoxs/api.simple-twitter/pkg/auth"
	"github.com/skyfoxs/api.simple-twitter/pkg/constants"
	"github.com/skyfoxs/api.simple-twitter/pkg/handler"
)

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
		uid, err := auth.ExtractUserIDFromToken(s[1], m.Secret)
		if err != nil {
			handler.Unauthorized(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), constants.UserID, uid)
		f.ServeHTTP(w, r.WithContext(ctx))
	}
}
