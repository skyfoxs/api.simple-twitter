package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var addr = ":8080"

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		logger: logger,
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Start server")
	err := server.ListenAndServe()
	logger.Fatal(err)
}
