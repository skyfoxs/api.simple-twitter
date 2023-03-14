package main

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	ID        string `json:"-"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Image     *Image `json:"-"`
}

type Image struct {
	Data []byte
	Type string
}

func (app *application) profileImage(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(ContextUserIDKey).(string)
	if getUserByID(id) == nil || getUserByID(id).Image == nil {
		app.respondNotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(getUserByID(id).Image.Data)
}

func (app *application) profile(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(ContextUserIDKey).(string)
	if getUserByID(id) == nil || getUserByID(id).Image == nil {
		app.respondNotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getUserByID(id))
}
