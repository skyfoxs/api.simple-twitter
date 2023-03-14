package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type application struct {
	logger *log.Logger
}

func (app *application) respondInternalError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorResponse{Message: "Internal server error"})
}

func (app *application) respondNotFound(w http.ResponseWriter, r *http.Request) {
	notFoundHandler{}.ServeHTTP(w, r)
}
