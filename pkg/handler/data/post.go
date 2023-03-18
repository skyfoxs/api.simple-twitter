package data

import (
	"time"

	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
)

type PostRequest struct {
	Message string `json:"message"`
}

type PostResponse struct {
	ID             string    `json:"id"`
	Message        string    `json:"message"`
	UserID         string    `json:"-"`
	Datetime       time.Time `json:"datetime"`
	ConversationID *string   `json:"-"`
	Likes          []string  `json:"-"`
}

func NewPostResponse(p *idata.Post) PostResponse {
	return PostResponse{
		ID:       p.ID,
		Message:  p.Message,
		Datetime: p.Datetime,
	}
}
