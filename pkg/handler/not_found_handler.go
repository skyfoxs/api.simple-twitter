package handler

import (
	"encoding/json"
	"net/http"

	"github.com/skyfoxs/api.simple-twitter/pkg/handler/data"
)

type NotFoundHandler struct{}

func (n NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(data.ErrorResponse{Message: "Resource not found"})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	NotFoundHandler{}.ServeHTTP(w, r)
}
