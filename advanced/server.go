package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var SQLFILE string = ""

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"user"`
	Password  string `json:"password"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin"`
	Active    int    `json:"active"`
}

type Input struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Admin    int    `json:"admin"`
}

type UserPass struct {
	Username string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type V2Input struct {
	Uusername string `json:"username"`
	Upassword string `json:"password"`
	U         User   `json:"load"`
}

func main() {
	fmt.Println("Hello!")
}
