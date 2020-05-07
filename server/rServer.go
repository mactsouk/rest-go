package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mactsouk/handlers"
)

// PORT is where the web server listens to
var PORT = ":1234"

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}
	fmt.Println("Listening to", PORT)

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

	mux.HandleFunc("/", handlers.DefaultHandler)

	// Register GET
	getMux := mux.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/time", handlers.TimeHandler)
	getMux.HandleFunc("/getall", handlers.GetAllHandler)
	getMux.HandleFunc("/get/{id:[0-9]+}", handlers.GetHandler)

	// Register PUT
	// Update User

	// Register POST
	// Change + Add User

	// Register DELETE
	// Delete User

	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
}
