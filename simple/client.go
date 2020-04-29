package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

var u1 = User{"admin", "admin"}
var u2 = User{"tsoukalos", "pass"}
var u3 = User{"", "pass"}

func addEndpoint(server string, user User) int {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewReader(userMarshall)

	req, err := http.NewRequest("POST", server+addEndPoint, u)
	if err != nil {
		fmt.Println("Error is req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		return resp.StatusCode
	}

	return resp.StatusCode
}

const addEndPoint = "/add"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Need: Server")
		return
	}

	server := os.Args[1]
	HTTPcode := addEndpoint(server, u1)

	if HTTPcode != http.StatusOK {
		fmt.Println("Return code:", HTTPcode)
		return
	} else {
		fmt.Println("Data added:", u1)
	}

	HTTPcode = addEndpoint(server, u3)

	if HTTPcode != http.StatusOK {
		fmt.Println("Return code:", HTTPcode)
		return
	} else {
		fmt.Println("Data added:", u3)
	}

}
