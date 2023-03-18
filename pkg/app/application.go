package app

import (
	"log"

	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

type Application struct {
	Logger    *log.Logger
	UserModel *model.UserModel
	PostModel *model.PostModel
	SecretKey []byte
}
