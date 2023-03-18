package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/skyfoxs/api.simple-twitter/pkg/constants"
	"github.com/skyfoxs/api.simple-twitter/pkg/handler/data"
	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

type PostHandler struct {
	Logger    *log.Logger
	PostModel model.PostModel
}

func (h PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(constants.UserID).(string)
	var pr data.PostRequest
	err := json.NewDecoder(r.Body).Decode(&pr)
	if err != nil {
		h.Logger.Printf("%v\n", err)
		BadRequest(w, r)
		return
	}

	p := idata.Post{
		ID:       uuid.NewString(),
		Message:  pr.Message,
		UserID:   uid,
		Datetime: time.Now(),
	}
	h.PostModel.Add(p)
	h.Logger.Printf("create post: %v %v user: %v\n", p.ID, p.Message, p.UserID)
	w.WriteHeader(http.StatusCreated)
}
