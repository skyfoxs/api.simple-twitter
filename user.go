package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

var u = []Profile{}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
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

	p := Profile{
		ID:        uuid.NewString(),
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Password:  md5str(r.FormValue("password")),
		Image:     image,
	}
	if getUserByEmail(p.Email) != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "This email already taken"})
		return
	}
	u = append(u, p)
	app.logger.Printf("create user: %v %v\n", p.ID, p.Email)
	w.WriteHeader(http.StatusCreated)
}

func getUserByEmail(e string) *Profile {
	for _, v := range u {
		if v.Email == e {
			return &v
		}
	}
	return nil
}

func getUserByID(id string) *Profile {
	for _, v := range u {
		if v.ID == id {
			return &v
		}
	}
	return nil
}

func md5str(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
