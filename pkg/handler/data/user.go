package data

import (
	"fmt"

	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
)

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

type FollowingResponse struct {
	Following []ProfileResponse `json:"following"`
}

func NewFollowingResponse(l []idata.Profile) FollowingResponse {
	f := []ProfileResponse{}
	for _, v := range l {
		f = append(f, NewProfileResponse(&v))
	}
	return FollowingResponse{
		Following: f,
	}
}

type GetPostResponse struct {
	Posts []PostResponse `json:"posts"`
}

func NewGetPostResponse(pl []idata.Post) GetPostResponse {
	result := []PostResponse{}
	for _, v := range pl {
		result = append(result, NewPostResponse(&v))
	}
	return GetPostResponse{
		Posts: result,
	}
}
