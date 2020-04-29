package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func getEndpoint(server string, user User) int {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewReader(userMarshall)

	req, err := http.NewRequest("GET", server+getEndPoint, u)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp == nil {
		return http.StatusNotFound
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("/get Response:", string(data))

	fmt.Println("/get Status:", resp.StatusCode)
	return resp.StatusCode
}

func addEndpoint(server string, user User) int {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewReader(userMarshall)

	req, err := http.NewRequest("POST", server+addEndPoint, u)
	if err != nil {
		fmt.Println("Error in req: ", err)
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
const getEndPoint = "/get"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Need: Server")
		return
	}
	server := os.Args[1]

	fmt.Println("/add")
	HTTPcode := addEndpoint(server, u1)
	if HTTPcode != http.StatusOK {
		fmt.Println("u1 Return code:", HTTPcode)
	} else {
		fmt.Println("u1 Data added:", u1, HTTPcode)
	}

	HTTPcode = addEndpoint(server, u2)
	if HTTPcode != http.StatusOK {
		fmt.Println("u2 Return code:", HTTPcode)
	} else {
		fmt.Println("u2 Data added:", u2, HTTPcode)
	}

	HTTPcode = addEndpoint(server, u3)
	if HTTPcode != http.StatusOK {
		fmt.Println("u3 Return code:", HTTPcode)
	} else {
		fmt.Println("u3 Data added:", u3, HTTPcode)
	}

	fmt.Println("/get")
	HTTPcode = getEndpoint(server, u1)
	fmt.Println("/get u1 return code:", HTTPcode)

	HTTPcode = getEndpoint(server, u2)
	fmt.Println("/get u2 return code:", HTTPcode)

	HTTPcode = getEndpoint(server, u3)
	fmt.Println("/get u3 return code:", HTTPcode)
}
