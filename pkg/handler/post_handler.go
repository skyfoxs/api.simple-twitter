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
	UserModel *model.UserModel
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
	if ok := h.PostModel.AddComment(pid, c); !ok {
		NotFound(w, r)
		return
	}
	h.Logger.Printf("comment to %v : %v %v user: %v\n", pid, c.ID, c.Message, c.UserID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data.NewPostResponse(&c))
}

func (h PostHandler) GetPostById(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")
	h.Logger.Printf("get post id: %v\n", id)
	p := h.PostModel.GetByID(id)
	if p == nil {
		NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.NewPostResponse(p))
}

func (h PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("type")
	uid := r.URL.Query().Get("userId")
	h.Logger.Printf("get post query params: type=%v, userId=%v\n", t, uid)
	if t == "individual" {
		h.Logger.Printf("get individual post with user id: %v\n", uid)
		p := h.UserModel.GetByID(uid)
		if p == nil {
			NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.NewGetPostResponse(h.PostModel.GetByUserID(uid)))
		return
	}
	if t == "feed" {
		id := r.Context().Value(constants.UserID).(string)
		p := h.UserModel.GetByID(id)
		if p == nil {
			NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.NewGetPostResponse(h.PostModel.GetFeed(id, p.Following)))
		return
	}
	BadRequest(w, r)
}
