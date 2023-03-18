package data

import (
	"fmt"

	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ProfileResponse struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	ImageURL  string `json:"imageURL,omitempty"`
}

func NewProfileResponse(p *idata.Profile) ProfileResponse {
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
