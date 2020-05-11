package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
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
	if len(arguments) >= 2 {
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
	getMux.HandleFunc("/getid", handlers.GetIDHandler)
	getMux.HandleFunc("/logged", handlers.LoggedUsersHandler)
	getMux.HandleFunc("/username/{id:[0-9]+}", handlers.GetUserDataHandler)

	// Register PUT
	// Update User
	putMux := mux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/update", handlers.UpdateHandler)

	// Register POST
	// Add User + Login + Logout
	postMux := mux.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/add", handlers.AddHandler)
	postMux.HandleFunc("/login", handlers.LoginHandler)
	postMux.HandleFunc("/logout", handlers.LogoutHandler)

	// Register DELETE
	// Delete User
	deleteMux := mux.Methods(http.MethodDelete).Subrouter()
	deleteMux.HandleFunc("/username/{id:[0-9]+}", handlers.DeleteHandler)
	deleteMux.HandleFunc("/", handlers.DefaultHandler)

	go func() {
		log.Println("Listening to", PORT)
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			return
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	sig := <-sigs
	log.Println("Quitting after signal:", sig)
	time.Sleep(5 * time.Second)
	s.Shutdown(nil)
}
