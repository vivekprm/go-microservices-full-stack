package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vivekprm/go-corelib/middlewares"
	"github.com/vivekprm/go-microservices-full-stack/todo-service/config"
	"github.com/vivekprm/go-microservices-full-stack/todo-service/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("unable to load env file: %v", err)
	}
	cfg := config.GetConfig()

	store, err := db.ConnectDB(cfg)

	todoMux := http.NewServeMux()

	corsConfig := middlewares.CorsConfig{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	}
	todoMux.Handle("/api/todos", &middlewares.CorsHandler{Next: &middlewares.JwtHandler{Next: &TodoHandler{db: store}}, Config: &corsConfig})
	todoMux.Handle("/api/todos/", &middlewares.CorsHandler{Next: &TodoHandler{}, Config: &corsConfig})

	s := http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: todoMux,
	}
	s.ListenAndServe()
}
