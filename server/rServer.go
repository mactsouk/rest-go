package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mactsouk/handlers"
)

// PORT is where the web server listens to
var PORT = ":1234"

type notAllowedHandler struct{}

func (h notAllowedHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	handlers.MethodNotAllowedHandler(rw, r)
}

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

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

	mux.NotFoundHandler = http.HandlerFunc(handlers.DefaultHandler)

	notAllowed := notAllowedHandler{}
	mux.MethodNotAllowedHandler = notAllowed

	// Register GET
	getMux := mux.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/time", handlers.TimeHandler)
	getMux.HandleFunc("/getall", handlers.GetAllHandler)
	getMux.HandleFunc("/get/{id:[0-9]+}", handlers.GetHandler)

	// Register PUT
	// Update User
	putMux := mux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/update", handlers.UpdateHandler)

	// Register POST
	// Change + Add User
	postMux := mux.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/add", handlers.AddHandler)
	postMux.HandleFunc("/login", handlers.LoginHandler)
	postMux.HandleFunc("/logout", handlers.LogoutHandler)

	// Register DELETE
	// Delete User
	deleteMux := mux.Methods(http.MethodDelete).Subrouter()
	deleteMux.HandleFunc("/delete", handlers.DeleteHandler)
	deleteMux.HandleFunc("/", handlers.DefaultHandler)

	log.Println("Listening to", PORT)

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
