package main

import (
	"encoding/json"
	"io"
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
	if getUserByID(id) == nil {
		app.respondNotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getUserByID(id))
}

func (app *application) patchProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(ContextUserIDKey).(string)
	u := getUserByID(id)
	if u == nil {
		app.respondNotFound(w, r)
		return
	}
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		app.logger.Printf("%v\n", err)
		app.respondInternalError(w, r)
		return
	}

	var image *Image = nil
	if _, ok := r.MultipartForm.File["image"]; ok {
		f, _, err := r.FormFile("image")
		if err != nil {
			app.logger.Printf("%v\n", err)
			app.respondInternalError(w, r)
			return
		}
		defer f.Close()

		b, err := io.ReadAll(f)
		if err != nil {
			app.logger.Printf("%v\n", err)
			app.respondInternalError(w, r)
			return
		}
		filetype := http.DetectContentType(b)
		app.logger.Printf("file type: %v\n", filetype)
		image = &Image{
			Data: b,
			Type: filetype,
		}
	}

	fn := r.FormValue("firstname")
	if fn != "" {
		u.Firstname = fn
	}
	ln := r.FormValue("lastname")
	if ln != "" {
		u.Lastname = ln
	}
	if image != nil {
		u.Image = image
	}
	app.logger.Printf("patch user: %v %v %v %v\n", u.ID, u.Email, u.Firstname, u.Lastname)
	w.WriteHeader(http.StatusAccepted)
}
