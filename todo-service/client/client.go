package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	createTodo()
	updateTodo()
	getTodo()
}

func getTodo() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/todos/1", nil)
	if err != nil {
		log.Printf("error in creating get todo request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		log.Printf("error in getting todo: %v\n", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error in reading response: %v\n", err)
	}
	log.Println(string(data))
}

func updateTodo() {
	req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/todos/1", nil)
	if err != nil {
		log.Printf("error in creating update todo request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		log.Printf("error in updating todo: %v\n", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error in reading response: %v\n", err)
	}
	log.Println(string(data))
}

func createTodo() {
	resp, err := http.Post("http://localhost:5000/api/todos", "application/json", bytes.NewBufferString(`
		{
			"name": "todo1",
			"description": "clean the kitchen",
			"status": "pending",
		}
	`))
	if err != nil {
		log.Printf("error in creating todo: %v\n", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error in reading response: %v\n", err)
	}
	log.Println(string(data))
}
