package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mactsouk/handlers"
)

var SQLFILE string = ""
var PORT string = ":1234"

func main() {
	log.Println("Hello!")

	mux := mux.NewRouter()
	s := http.Server{
		Addr:         PORT,
		Handler:      mux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	putMux := mux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/v2/files/{id:[0-9]+}", handlers.UploadFile)

	getMux := mux.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/v2/files/{id:[0-9]+}", handlers.SendFile)

	log.Println("Listening to", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
