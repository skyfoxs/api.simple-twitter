package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/skyfoxs/api.simple-twitter/pkg/auth"
	"github.com/skyfoxs/api.simple-twitter/pkg/handler/data"
	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

type AuthHandler struct {
	Logger    *log.Logger
	UserModel *model.UserModel
	SecretKey []byte
}

func (h AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var c data.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		h.Logger.Printf("%v\n", err)
		BadRequest(w, r)
		return
	}
	if p, ok := h.UserModel.Get(c.Email, c.Password); ok {
		t, err := auth.NewToken(p, h.SecretKey)
		if err != nil {
			h.Logger.Printf("%v\n", err)
			InternalServerError(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.SignInResponse{
			UserID: p.ID,
			Token:  t,
		})
		return
	}
	Unauthorized(w, r)
}
