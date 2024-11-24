package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vivekprm/go-corelib/middlewares"
	"github.com/vivekprm/go-microservices-full-stack/user-service/config"
	"github.com/vivekprm/go-microservices-full-stack/user-service/db"
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
	userMux.Handle("/api/users/", &middlewares.JwtHandler{Next: &userHandler{db: db}})
	userMux.Handle("/api/login", &middlewares.CorsHandler{
		Next: &userHandler{db: db},
		Config: &middlewares.CorsConfig{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"*"},
			AllowedHeaders: []string{"*"},
		},
	})
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: userMux,
	}
	s.ListenAndServe()
}
