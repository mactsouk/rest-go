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

	mux := mux.NewRouter()

	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	fmt.Println(s)

	record := handlers.User{}
	fmt.Println(record)

	fmt.Println("Hello!")
}
