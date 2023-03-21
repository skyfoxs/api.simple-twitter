package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/skyfoxs/api.simple-twitter/pkg/constants"
	"github.com/skyfoxs/api.simple-twitter/pkg/handler/data"
	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

type PostHandler struct {
	Logger    *log.Logger
	PostModel *model.PostModel
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data.NewPostResponse(&p))
}

func (h PostHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	pid := httprouter.ParamsFromContext(r.Context()).ByName("id")
	uid := r.Context().Value(constants.UserID).(string)
	var pr data.PostRequest
	err := json.NewDecoder(r.Body).Decode(&pr)
	if err != nil {
		h.Logger.Printf("%v\n", err)
		BadRequest(w, r)
		return
	}

	c := idata.Post{
		ID:       uuid.NewString(),
		Message:  pr.Message,
		UserID:   uid,
		Datetime: time.Now(),
	}
	p := h.PostModel.GetByID(pid)
	if p.Comments == nil {
		p.Comments = []idata.Post{c}
	} else {
		p.Comments = append(p.Comments, c)
	}
	h.Logger.Printf("comment to %v : %v %v user: %v\n", pid, c.ID, c.Message, c.UserID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data.NewPostResponse(&c))
}
