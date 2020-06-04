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
var IMAGESPATH = "/tmp/files"

type notAllowedHandler struct{}

func (h notAllowedHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	handlers.MethodNotAllowedHandler(rw, r)
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		log.Println("Not enough arguments: SQLFILE IMAGESPATH [PORT]")
		return
	}

	if len(arguments) == 2 {
		SQLFILE = arguments[1]
	} else if len(arguments) == 3 {
		SQLFILE = arguments[1]
		IMAGESPATH = arguments[2]
	} else if len(arguments) == 4 {
		SQLFILE = arguments[1]
		IMAGESPATH = arguments[2]
		PORT = ":" + arguments[3]
	}

	handlers.IMAGESPATH = IMAGESPATH
	err := handlers.CreateImageDirectory()
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

	getMux.HandleFunc("/v1/time", handlers.TimeHandler)
	getMux.HandleFunc("/v1/getall", handlers.GetAllHandlerUpdated)
	getMux.HandleFunc("/v2/time", handlers.TimeHandler)
	getMux.HandleFunc("/v2/getall", handlers.GetAllHandlerV2)

	postMux := mux.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/v1/add", handlers.AddHandler)
	postMux.HandleFunc("/v1/login", handlers.LoginHandler)
	postMux.HandleFunc("/v1/logout", handlers.LogoutHandler)
	postMux.HandleFunc("/v2/add", handlers.AddHandlerV2)
	postMux.HandleFunc("/v2/login", handlers.LoginHandlerV2)
	postMux.HandleFunc("/v2/logout", handlers.LogoutHandlerV2)

	mux.Use(handlers.MiddleWare)

	log.Println("Listening to", PORT)
	err = s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
