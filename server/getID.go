package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
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
	log.Println("ID Serving:", r.URL.Path, "from", r.Host)
	reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)
	if len(g) != 1 {
		log.Println("Invalid URI - More than one ID!")
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}

	idText := g[0][1]
	id, err := strconv.Atoi(idText)
	if err != nil {
		log.Println("Cannot convert to integer", idText)
		http.Error(w, "Invalid URI", http.StatusBadRequest)
		return
	}

	log.Println("ID:", id)
	Body := "Looking for information about ID: " + idText
	fmt.Fprintf(w, "%s\n", Body)
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
	mux.Handle("/id/", http.HandlerFunc(idHandler))

	fmt.Println("Ready to serve at", PORT)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

}
