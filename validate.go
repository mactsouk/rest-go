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
	ID        int    `json:"id" validate:"gte=1"`
	Username  string `json:"user" validate:"required"`
	Password  string `json:"password" validate:"required"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin" validate:"required"`
	Active    int    `json:"active"`
}

type U3 struct {
	ID        int    `json:"id" validate:"required,gte=1"`
	Username  string `json:"user" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	LastLogin int64  `json:"lastlogin" validate:"gte=1590550000,,lte=1800000000"`
	Admin     int    `json:"admin" validate:"required"`
	Active    int    `json:"active"`
}

var validate *validator.Validate

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a JSON record as input!")
		return
	}
	input := arguments[1]
	buf := []byte(input)

	validate = validator.New()

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
	var u2 U2
	err = json.Unmarshal(buf, &u2)
	if err != nil {
		fmt.Println("U1 Unmarshal:", err)
	}
	fmt.Println("U2:", u2)

	err = u2.Validate()
	if err != nil {
		fmt.Println("U2 - Validate:", err)
	}

	// U3
	var u3 U3
	err = json.Unmarshal(buf, &u3)
	if err != nil {
		fmt.Println("U3 Unmarshal:", err)
	}
	fmt.Println("U3:", u3)

	err = u3.Validate()
	if err != nil {
		fmt.Println("U3 - Validate:", err)
	}
}

// Validate method validates the data of UserPass
func (p *U1) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

func (p *U2) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

func (p *U3) Validate() error {
	validate := validator.New()
	validate.RegisterStructValidation(UserPasswordLength, U3{})
	return validate.Struct(p)
}

func UserPasswordLength(sl validator.StructLevel) {
	user := sl.Current().Interface().(U3)
	if len(user.Username) < 10 || len(user.Password) < 8 {
		sl.ReportError(user.Username, "user", "Username", "usernameOrPassword", "")
		sl.ReportError(user.Password, "password", "Password", "usernameOrPassword", "")
	}
}
