package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var secretKey = []byte("SecretYouShouldHide")

type TokenClaims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

type Authentication struct {
	Token string `json:"token"`
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var c Credential
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		app.logger.Printf("%v\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Bad request"})
		return
	}
	u := getUserByEmail(c.Email)

	if u == nil || u.Password != md5str(c.Password) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid credentials"})
		return
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		u.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "simple-twitter",
		},
	})

	token, err := t.SignedString(secretKey)
	if err != nil {
		app.logger.Printf("%v\n", err)
		app.respondInternalError(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Authentication{Token: token})
}
