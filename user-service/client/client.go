package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	createUser()
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

func updateUser() {
	req, err := http.NewRequest(http.MethodPut, "http://localhost:4000/api/users/1", bytes.NewBufferString(`
	{
		"firstName": "vivek",
		"lastName": "mishra",
		"email": "vivek@xyz.com"
	}
	`))
	req.Header.Set("Content-Type", "application/json")
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
