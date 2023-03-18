package handler

import (
	"encoding/json"
	"net/http"

	"github.com/skyfoxs/api.simple-twitter/pkg/handler/data"
)

type InternalServerErrorHandler struct{}

func (h InternalServerErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(data.ErrorResponse{Message: "Internal server error"})
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	InternalServerErrorHandler{}.ServeHTTP(w, r)
}
