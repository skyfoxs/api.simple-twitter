package idata

import (
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/skyfoxs/api.simple-twitter/pkg/encrypt"
)

type Profile struct {
	ID        string
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Image     *Image
	Following []string
}

type Image struct {
	Data []byte
	Type string
}

func NewProfileFromMultipartFormData(r *http.Request) (*Profile, error) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil, err
	}

	var image *Image = nil
	if _, ok := r.MultipartForm.File["image"]; ok {
		f, _, err := r.FormFile("image")
		if err != nil {
			return nil, err
		}
		defer f.Close()

		b, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}
		filetype := http.DetectContentType(b)
		image = &Image{
			Data: b,
			Type: filetype,
		}
	}

	return &Profile{
		ID:        uuid.NewString(),
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Password:  encrypt.MD5str(r.FormValue("password")),
		Image:     image,
	}, nil
}
