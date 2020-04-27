package main

import (
	"fmt"
	"net/http"
	"time"
)

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
