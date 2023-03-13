package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

var image []byte

type application struct {
	logger *log.Logger
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it's works"))
}

func (app *application) uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	f, fh, err := r.FormFile("image")
	if err != nil {
		app.logger.Printf("%v\n", err)
		return
	}
	defer f.Close()
	app.logger.Printf("%v\n", fh.Header)

	fb, err := io.ReadAll(f)
	if err != nil {
		app.logger.Printf("%v\n", err)
		return
	}
	image = fb
}

func (app *application) downloadImageHandler(w http.ResponseWriter, r *http.Request) {
	if image == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(image)
}

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/image", app.downloadImageHandler)
	router.HandlerFunc(http.MethodPost, "/image", app.uploadImageHandler)
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
