package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

type application struct {
	logger *log.Logger
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it's works"))
}

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", app.home)
	return router
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		logger: logger,
	}

	server := &http.Server{
		Addr:         ":8080",
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Start server")
	err := server.ListenAndServe()
	logger.Fatal(err)
}
