package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var PORT = ":1234"

func timeHandler(rw http.ResponseWriter, r *http.Request) {

}

func addHandler(rw http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := mux.NewRouter()

	putMux := mux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/time", timeHandler)

	getMux := mux.Methods(http.MethodGet).Subrouter()
	getMux.Handle("/add", addHandler)

	s := http.Server{
		Addr:         PORT,
		Handler:      mux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Println("Listening to", PORT)
	err = s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
