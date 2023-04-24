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
	user := handler.UserHandler{
		Logger:    app.Logger,
		UserModel: app.UserModel,
	}
	post := handler.PostHandler{
		Logger:    app.Logger,
		UserModel: app.UserModel,
		PostModel: app.PostModel,
	}

	r := httprouter.New()
	r.NotFound = handler.NotFoundHandler{}

	r.HandlerFunc(http.MethodGet, "/posts", m.TokenRequired(post.GetPosts))
	r.HandlerFunc(http.MethodGet, "/posts/:id", m.TokenRequired(post.GetPostById))
	r.HandlerFunc(http.MethodGet, "/users", m.TokenRequired(user.Search))
	r.HandlerFunc(http.MethodGet, "/users/:id", m.TokenRequired(user.Info))
	r.HandlerFunc(http.MethodGet, "/users/:id/image", m.TokenRequired(user.Image))
	r.HandlerFunc(http.MethodGet, "/users/:id/following", m.TokenRequired(user.GetFollowing))

	r.HandlerFunc(http.MethodPost, "/posts", m.TokenRequired(post.Create))
	r.HandlerFunc(http.MethodPost, "/posts/:id/comment", m.TokenRequired(post.CreateComment))
	r.HandlerFunc(http.MethodPost, "/users/:id/following", m.TokenRequired(user.AddFollowing))

	r.HandlerFunc(http.MethodDelete, "/users/:id/following", m.TokenRequired(user.DeleteFollowing))

	r.HandlerFunc(http.MethodPatch, "/users/:id", m.TokenRequired(user.Patch))

	r.HandlerFunc(http.MethodPost, "/users", user.Create)
	r.HandlerFunc(http.MethodPost, "/signin", a.SignIn)
	return r
}
