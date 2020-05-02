package main

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"user"`
	Password  string `json:"password"`
	LastLogin time   `json:"lastlogin"`
	Admin     bool   `json:"admin"`
	Active    bool   `json:"active"`
}

func main() {

	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	fmt.Println("Hello!")
}
