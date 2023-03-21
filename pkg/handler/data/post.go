package data

import (
	"time"

	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
)

type PostRequest struct {
	Message string `json:"message"`
}

type PostResponse struct {
	ID       string         `json:"id"`
	Message  string         `json:"message"`
	UserID   string         `json:"userId"`
	Datetime time.Time      `json:"datetime"`
	Comments []PostResponse `json:"comments"`
}

func NewPostResponse(p *idata.Post) PostResponse {
	c := []PostResponse{}
	if p.Comments != nil {
		for _, v := range p.Comments {
			c = append(c, NewPostResponse(&v))
		}
	}
	return PostResponse{
		ID:       p.ID,
		UserID:   p.UserID,
		Message:  p.Message,
		Datetime: p.Datetime,
		Comments: c,
	}
}
