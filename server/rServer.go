package main

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"user"`
	Password string `json:"password"`
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
