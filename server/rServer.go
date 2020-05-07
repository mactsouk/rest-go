package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mactsouk/handlers"
)

var PORT = ":1234"

func main() {

	// Create a new ServeMux using Gorilla
	mux := mux.NewRouter()

	s := http.Server{
		Addr:         PORT,
		Handler:      mux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	mux.HandleFunc("/", handlers.DefaultHandler())
	mux.HandleFunc("/time", handlers.TimeHandler())

	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
}
