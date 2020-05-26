package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

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

func main() {
	fmt.Println("Hello!")
	var u UserPass
	err := u.Validate()
	if err != nil {
		fmt.Println("IsUserAdmin - Validate:", err)
		return false
	}
}

// Validate method validates the data of UserPass
func (p *UserPass) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
