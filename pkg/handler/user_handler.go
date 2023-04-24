package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skyfoxs/api.simple-twitter/pkg/constants"
	"github.com/skyfoxs/api.simple-twitter/pkg/handler/data"
	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

type UserHandler struct {
	Logger    *log.Logger
	UserModel *model.UserModel
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
	json.NewEncoder(w).Encode(data.NewFollowingResponse(h.UserModel.GetFollowing(id)))
}

func (h UserHandler) Patch(w http.ResponseWriter, r *http.Request) {
	uid := httprouter.ParamsFromContext(r.Context()).ByName("id")
	id := r.Context().Value(constants.UserID).(string)
	if uid != id {
		Unauthorized(w, r)
		return
	}
	p := h.UserModel.GetByID(id)
	if p == nil {
		NotFound(w, r)
		return
	}
	t, err := idata.NewProfileFromMultipartFormData(r)
	if err != nil {
		h.Logger.Printf("%v\n", err)
		InternalServerError(w, r)
		return
	}
	fn := t.Firstname
	if fn != "" {
		p.Firstname = fn
	}
	ln := t.Lastname
	if ln != "" {
		p.Lastname = ln
	}
	if t.Image != nil {
		p.Image = t.Image
	}
	h.Logger.Printf("patch user: %v %v %v %v\n", p.ID, p.Email, p.Firstname, p.Lastname)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.NewProfileResponse(p))
}

func (h UserHandler) AddFollowing(w http.ResponseWriter, r *http.Request) {
	uid := httprouter.ParamsFromContext(r.Context()).ByName("id")
	id := r.Context().Value(constants.UserID).(string)
	if uid != id {
		Unauthorized(w, r)
		return
	}
	var req data.FollowingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.Logger.Printf("%v\n", err)
		BadRequest(w, r)
		return
	}
	if ok := h.UserModel.AddFollowing(id, req.ID); !ok {
		BadRequest(w, r)
		return
	}
	h.Logger.Printf("user %v start follow user %v\n", id, req.ID)
	w.WriteHeader(http.StatusNoContent)
}

func (h UserHandler) DeleteFollowing(w http.ResponseWriter, r *http.Request) {
	uid := httprouter.ParamsFromContext(r.Context()).ByName("id")
	id := r.Context().Value(constants.UserID).(string)
	if uid != id {
		Unauthorized(w, r)
		return
	}
	var req data.FollowingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.Logger.Printf("%v\n", err)
		BadRequest(w, r)
		return
	}
	if ok := h.UserModel.DeleteFollowing(id, req.ID); !ok {
		BadRequest(w, r)
		return
	}
	h.Logger.Printf("user %v unfollow user %v\n", id, req.ID)
	w.WriteHeader(http.StatusNoContent)
}
