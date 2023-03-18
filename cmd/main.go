package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/skyfoxs/api.simple-twitter/pkg/app"
	"github.com/skyfoxs/api.simple-twitter/pkg/model"
)

var addr = ":8080"

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	a := &app.Application{
		Logger:    logger,
		UserModel: model.NewUserModel(),
		PostModel: model.NewPostModel(),
		SecretKey: []byte("SecretYouShouldHide"),
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      a.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Start server")
	err := server.ListenAndServe()
	logger.Fatal(err)
}
