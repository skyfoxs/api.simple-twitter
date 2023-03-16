package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

type UnauthorizedHandler struct{}

func (h UnauthorizedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(errors.New("Unauthorized access"))
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	UnauthorizedHandler{}.ServeHTTP(w, r)
}
