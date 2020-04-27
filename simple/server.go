package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var PORT = ":1234"

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

var user User

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving:", r.URL.Path, "from", r.Host)
	log.Printf("Data: %s\n", "Bye!")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:" + t + "\n"
	fmt.Fprintf(w, "%s", Body)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	log.Println(user.Username)
	log.Println(user.Password)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	d, err := json.Marshal(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s\n", d)

}

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
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

	mux.Handle("/time", http.HandlerFunc(timeHandler))
	mux.Handle("/add", http.HandlerFunc(addHandler))
	mux.Handle("/get", http.HandlerFunc(getHandler))
	mux.Handle("/", http.HandlerFunc(defaultHandler))

	fmt.Println("Ready to serve at", PORT)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
