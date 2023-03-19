package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skyfoxs/api.simple-twitter/pkg/handler"
	"github.com/skyfoxs/api.simple-twitter/pkg/middleware"
)

func (app *Application) Routes() *httprouter.Router {
	m := middleware.JWTAuth{Secret: app.SecretKey}
	a := handler.AuthHandler{
		Logger:    app.Logger,
		UserModel: app.UserModel,
		SecretKey: app.SecretKey,
	}
	u := handler.UserHandler{
		Logger:    app.Logger,
		UserModel: app.UserModel,
	}
	p := handler.ProfileHandler{
		Logger:    app.Logger,
		UserModel: app.UserModel,
	}
	po := handler.PostHandler{
		Logger:    app.Logger,
		PostModel: *app.PostModel,
	}

	r := httprouter.New()
	r.NotFound = handler.NotFoundHandler{}

	r.HandlerFunc(http.MethodGet, "/user/:id", u.Info)
	r.HandlerFunc(http.MethodGet, "/user/:id/image", u.Image)
	r.HandlerFunc(http.MethodGet, "/user/:id/following", (u.GetFollowing))
	r.HandlerFunc(http.MethodPost, "/user", u.Create)

	r.HandlerFunc(http.MethodPost, "/login", a.Login)

	r.HandlerFunc(http.MethodGet, "/profile", m.TokenRequired(p.Info))
	r.HandlerFunc(http.MethodGet, "/profile/image", m.TokenRequired(p.Image))
	r.HandlerFunc(http.MethodGet, "/profile/following", m.TokenRequired(p.GetFollowing))
	r.HandlerFunc(http.MethodPatch, "/profile", m.TokenRequired(p.Patch))
	r.HandlerFunc(http.MethodPost, "/profile/following", m.TokenRequired(p.AddFollowing))
	r.HandlerFunc(http.MethodDelete, "/profile/following", m.TokenRequired(p.DeleteFollowing))

	r.HandlerFunc(http.MethodPost, "/post", m.TokenRequired(po.Create))
	return r
}
