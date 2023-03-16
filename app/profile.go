package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/skyfoxs/api.simple-twitter/pkg/handler"
	"github.com/skyfoxs/api.simple-twitter/pkg/middleware"
)

type Profile struct {
	ID        string `json:"id"`
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

type ProfileResponse struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	ImageURL  string `json:"imageURL,omitempty"`
}

func NewProfileResponse(p *Profile) ProfileResponse {
	var imgURL string
	if p.Image != nil {
		imgURL = fmt.Sprintf("/user/%s/image", p.ID)
	}
	return ProfileResponse{
		ID:        p.ID,
		Firstname: p.Firstname,
		Lastname:  p.Lastname,
		Email:     p.Email,
		ImageURL:  imgURL,
	}
}

func (app *Application) profileImage(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.UserID).(string)
	if getUserByID(id) == nil || getUserByID(id).Image == nil {
		handler.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(getUserByID(id).Image.Data)
}

func (app *Application) profile(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.UserID).(string)
	if getUserByID(id) == nil {
		handler.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(NewProfileResponse(getUserByID(id)))
}

func (app *Application) patchProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.UserID).(string)
	u := getUserByID(id)
	if u == nil {
		handler.NotFound(w, r)
		return
	}
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		app.Logger.Printf("%v\n", err)
		handler.InternalServerError(w, r)
		return
	}

	var image *Image = nil
	if _, ok := r.MultipartForm.File["image"]; ok {
		f, _, err := r.FormFile("image")
		if err != nil {
			app.Logger.Printf("%v\n", err)
			handler.InternalServerError(w, r)
			return
		}
		defer f.Close()

		b, err := io.ReadAll(f)
		if err != nil {
			app.Logger.Printf("%v\n", err)
			handler.InternalServerError(w, r)
			return
		}
		filetype := http.DetectContentType(b)
		app.Logger.Printf("file type: %v\n", filetype)
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
	app.Logger.Printf("patch user: %v %v %v %v\n", u.ID, u.Email, u.Firstname, u.Lastname)
	w.WriteHeader(http.StatusAccepted)
}
