package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type notFoundHandler struct{}

func (n notFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{Message: "Resource not found"})
}

func (app *application) routes() *httprouter.Router {
	r := httprouter.New()
	r.NotFound = notFoundHandler{}
	r.HandlerFunc(http.MethodPost, "/user", app.createUser)
	r.HandlerFunc(http.MethodGet, "/profile/:id/image", app.profileImage)
	return r
}
