package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var PORT = ":1234"

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving:", r.URL.Path, "from", r.Host)
	log.Printf("Data: %s\n", d)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:" + t
	fmt.Fprintf(w, "%s", Body)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
}

func main() {

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/", defaultHandler)

	fmt.Println("Ready to serve at", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
