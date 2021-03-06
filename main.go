package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rakhmadbudiono/simple-api/api"
)

func main() {
	startServer()
}

func startServer() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Handler:      api.Router(),
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Starting server on port %s!", port)

	log.Fatal(srv.ListenAndServe())
}
