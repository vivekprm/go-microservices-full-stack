package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	userMux := http.NewServeMux()
	userMux.Handle("/api/users/", &userHandler{})
	s := &http.Server{
		Addr:    ":4000",
		Handler: userMux,
	}
	go func() {
		log.Fatalln(s.ListenAndServe())
	}()
	fmt.Println("User service started. Press <ENTER> to stop.")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("User service stopped")
}
