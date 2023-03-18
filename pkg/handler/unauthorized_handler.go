package handler

import (
	"encoding/json"
	"net/http"

	"github.com/skyfoxs/api.simple-twitter/pkg/handler/data"
)

type UnauthorizedHandler struct{}

func (h UnauthorizedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(data.ErrorResponse{Message: "Unauthorized access"})
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	UnauthorizedHandler{}.ServeHTTP(w, r)
}
