package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/validator"
)

type U1 struct {
	ID        int    `json:"id"`
	Username  string `json:"user" validate:"required"`
	Password  string `json:"password" validate:"required"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin"`
	Active    int    `json:"active"`
}

type U2 struct {
	ID        int    `json:"id"`
	Username  string `json:"user" validate:"required"`
	Password  string `json:"password" validate:"required"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin" validate:"required"`
	Active    int    `json:"active"`
}

type U3 struct {
	ID        int    `json:"id" validate:"required"`
	Username  string `json:"user" validate:"required"`
	Password  string `json:"password" validate:"required"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin" validate:"required"`
	Active    int    `json:"active"`
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a JSON record as input!")
		return
	}
	input := arguments[1]
	buf := []byte(input)

	// U1
	var u1 U1

	err := json.Unmarshal(buf, &u1)
	if err != nil {
		fmt.Println("U1 Unmarshal:", err)
	}
	fmt.Println("U1:", u1)

	err = u1.Validate()
	if err != nil {
		fmt.Println("U1 - Validate:", err)
	}

	// U2

	// U3

}

// Validate method validates the data of UserPass
func (p *U1) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
