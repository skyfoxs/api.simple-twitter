package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skyfoxs/api.simple-twitter/pkg/handler/data"
	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

type UserHandler struct {
	Logger    *log.Logger
	UserModel *model.UserModel
	PostModel *model.PostModel
}

func (h UserHandler) Info(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")
	p := h.UserModel.GetByID(id)
	if p == nil {
		NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.NewProfileResponse(p))
}

func (h UserHandler) Image(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")
	p := h.UserModel.GetByID(id)
	if p == nil || p.Image == nil {
		NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", p.Image.Type)
	w.WriteHeader(http.StatusOK)
	w.Write(p.Image.Data)
}

func (h UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	p, err := idata.NewProfileFromMultipartFormData(r)
	if err != nil {
		h.Logger.Printf("%v\n", err)
		InternalServerError(w, r)
		return
	}
	if h.UserModel.GetByEmail(p.Email) != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(data.ErrorResponse{Message: "This email already taken"})
		return
	}
	h.UserModel.Add(*p)
	h.Logger.Printf("create user: %v %v\n", p.ID, p.Email)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data.NewProfileResponse(p))
}

func (h UserHandler) GetFollowing(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")
	p := h.UserModel.GetByID(id)
	if p == nil {
		NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.NewFollowingResponse(h.UserModel.GetFollowing(p)))
}

func (h UserHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")
	p := h.UserModel.GetByID(id)
	if p == nil {
		NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.NewGetPostResponse(h.PostModel.GetByUserID(id)))
}
