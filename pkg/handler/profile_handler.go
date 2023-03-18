package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/skyfoxs/api.simple-twitter/pkg/constants"
	"github.com/skyfoxs/api.simple-twitter/pkg/handler/data"
	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

type ProfileHandler struct {
	Logger    *log.Logger
	UserModel *model.UserModel
}

func (h ProfileHandler) Info(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(constants.UserID).(string)
	p := h.UserModel.GetByID(id)
	if p == nil {
		NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.NewProfileResponse(p))
}

func (h ProfileHandler) Image(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(constants.UserID).(string)
	p := h.UserModel.GetByID(id)
	if p == nil || p.Image == nil {
		NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", p.Image.Type)
	w.WriteHeader(http.StatusOK)
	w.Write(p.Image.Data)
}

func (h ProfileHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(constants.UserID).(string)
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
	w.WriteHeader(http.StatusAccepted)
}
