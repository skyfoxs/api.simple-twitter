package handler

import (
	"encoding/json"
	"net/http"

	"github.com/skyfoxs/api.simple-twitter/pkg/handler/data"
)

type BadRequestHandler struct{}

func (h BadRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(data.ErrorResponse{Message: "Bad request"})
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	BadRequestHandler{}.ServeHTTP(w, r)
}
