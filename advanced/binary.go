package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var PORT = ":1234"

func uploadFile(rw http.ResponseWriter, r *http.Request) {

}

func sendFile(rw http.ResponseWriter, r *http.Request) {

}

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

	// Register PUT
	// Update User
	putMux := mux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/upload/{id:[0-9]+}", uploadFile)

	getMux := mux.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/upload/{id:[0-9]+}", sendFile)

	log.Println("Listening to", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
