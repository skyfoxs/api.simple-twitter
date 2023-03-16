package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skyfoxs/api.simple-twitter/pkg/handler"
	"github.com/skyfoxs/api.simple-twitter/pkg/middleware"
)

func (app *Application) Routes() *httprouter.Router {
	r := httprouter.New()
	m := middleware.JWTAuth{Secret: app.SecretKey}
	r.NotFound = handler.NotFoundHandler{}

	r.HandlerFunc(http.MethodGet, "/user/:id", app.user)
	r.HandlerFunc(http.MethodGet, "/user/:id/image", app.userImage)
	r.HandlerFunc(http.MethodPost, "/user", app.createUser)

	r.HandlerFunc(http.MethodPost, "/login", app.login)

	r.HandlerFunc(http.MethodGet, "/profile", m.TokenRequired(app.profile))
	r.HandlerFunc(http.MethodGet, "/profile/image", m.TokenRequired(app.profileImage))
	r.HandlerFunc(http.MethodPatch, "/profile", m.TokenRequired(app.patchProfile))
	return r
}
