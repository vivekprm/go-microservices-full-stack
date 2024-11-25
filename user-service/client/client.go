package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/vivekprm/go-microservices-full-stack/user-service/models"
)

type LoginResponse struct {
	User  models.User
	Token string
}

func main() {
	// resp := login()
	// updateUser(resp.Token)
	createUser()
}
func login() LoginResponse {
	res, err := http.Post("http://localhost:4000/api/login", "application/json", bytes.NewBufferString(`
		{
			"email": "vivek@xyz.com",
			"password": "welcome"
		}
	`))
	if err != nil {
		log.Println(err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println("Response: ", string(data))
	var loginResponse LoginResponse
	json.Unmarshal(data, &loginResponse)
	return loginResponse
}
func createUser() {
	res, err := http.Post("http://localhost:4000/api/users", "application/json", bytes.NewBufferString(`
		{
			"firstName": "vivek",
			"lastName": "mishra",
			"email": "vivek@abc.com",
			"password": "welcome"
		}
	`))
	if err != nil {
		log.Println(err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println("Response: ", string(data))
}

func updateUser(token string) {
	req, err := http.NewRequest(http.MethodPut, "http://localhost:4000/api/users/1", bytes.NewBufferString(`
	{
		"firstName": "vivek",
		"lastName": "mishra",
		"email": "vivek@xyz.com"
	}
	`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("x-access-token", token)
	if err != nil {
		log.Fatal(err)
	}
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}
