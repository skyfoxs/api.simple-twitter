package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/skyfoxs/api.simple-twitter/pkg/idata"
)

func NewToken(u *idata.Profile, s []byte) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		u.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "simple-twitter",
		},
	})
	return t.SignedString(s)
}

func ExtractUserIDFromToken(t string, s []byte) (uid string, err error) {
	j, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) { return s, nil })
	if err != nil {
		return "", err
	}
	if claims, ok := j.Claims.(jwt.MapClaims); ok && j.Valid {
		return claims["userId"].(string), nil
	}
	return "", nil
}
