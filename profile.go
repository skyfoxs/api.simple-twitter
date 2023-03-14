package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Profile struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Image     *Image
}

type Image struct {
	Data []byte
	Type string
}

func (app *application) profileImage(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")
	if getUserByID(id) == nil || getUserByID(id).Image == nil {
		app.respondNotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(getUserByID(id).Image.Data)
}
