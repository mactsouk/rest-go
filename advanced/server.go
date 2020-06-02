package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

	log.Println("Listening to", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
