package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mactsouk/handlers"
)

var SQLFILE string = ""
var PORT string = ":1234"

type notAllowedHandler struct{}

func (h notAllowedHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	handlers.MethodNotAllowedHandler(rw, r)
}

func main() {
	arguments := os.Args
	if len(arguments) == 2 || len(arguments) == 3 {
		log.Println("Not enough arguments: DBNAME ")
		return
	}

	if len(arguments) == 5 {
		PORT = ":" + arguments[1]
	}

	err := handlers.CreateImageDirectory(IMAGESPATH)
	if err != nil {
		log.Println(err)
		return
	}

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

	putMux := mux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/v2/files/{filename:[a-zA-Z0-9][a-zA-Z0-9\\.]*[a-zA-Z0-9]}", handlers.UploadFile)

	getMux := mux.Methods(http.MethodGet).Subrouter()
	getMux.Handle(
		"/v2/files/{filename:[a-zA-Z0-9][a-zA-Z0-9\\.]*[a-zA-Z0-9]}",
		http.StripPrefix("/v2/files/", http.FileServer(http.Dir(IMAGESPATH))))

	mux.Use(handers.MiddleWare)

	log.Println("Listening to", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
