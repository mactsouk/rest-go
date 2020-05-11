package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// PORT is where the web server listens to
var PORT = ":1234"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

func idHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

func main() {
	arguments := os.Args
	if len(arguments) >= 2 {
		PORT = ":" + arguments[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/", http.HandlerFunc(defaultHandler))
	mux.Handle("/id", http.HandlerFunc(idHandler))

	fmt.Println("Ready to serve at", PORT)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

}
