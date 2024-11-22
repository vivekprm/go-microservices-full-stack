package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"user-service/config"
	"user-service/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment variables:", err)
	}

	// Connect to PostgreSQL database
	cfg := config.GetConfig()
	db, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	userMux := http.NewServeMux()
	userMux.Handle("/api/users", &userHandler{db: db})
	userMux.Handle("/api/users/", &userHandler{db: db})
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: userMux,
	}
	s.ListenAndServe()
}
