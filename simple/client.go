package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

func loginEndpoint(server string, user User) (int, string) {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewReader(userMarshall)

	req, err := http.NewRequest("GET", server+endPoint, u)
	if err != nil {
		fmt.Println("Error is req: ", err)
		return 400, ""
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		return resp.StatusCode, ""
	}

	data, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, string(data)
}

const endPoint = "/get"

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Need: Server username password")
		return
	}

	server := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	loginInfo := User{username, password}

	HTTPcode, data := loginEndpoint(server, loginInfo)

	if HTTPcode != 200 {
		fmt.Println("Return code:", HTTPcode)
		return
	}

	var user User
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(user.Username)
	log.Println(user.Password)
}
