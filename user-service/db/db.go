package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"user-service/config"
	"user-service/models"
	"golang.org/x/crypto/bcrypt"
)

type DB struct {
	*sql.DB
}

func ConnectDB(cfg *config.Config) (*DB, error) {
	connStr, err := cfg.GetDbConnectionString()
	if err != nil {
		return nil, err
	}
	log.Printf("Connection string: %s, driver name: %s", connStr, cfg.DriverName)
	db, err := sql.Open(cfg.DriverName, connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	// Check the database connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}
	log.Println("Database connected.")
	return &DB{db}, nil
}

func (db *DB) GetUsers() ([]models.User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("error getting users: %v", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *DB) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("error getting user by ID: %v", err)
	}

	return &user, nil
}

func (db *DB) AddUser(user *models.User) (models.User, error) {
	password := []byte(user.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("error creating hashed password: %v", err)
	}
	var id int
	err = db.QueryRow("INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id",
		user.FirstName, user.LastName, user.Email, hashedPassword).Scan(&id)
	if err != nil {
		return models.User{}, fmt.Errorf("error adding user: %v", err)
	}
	if err != nil {
		return models.User{}, fmt.Errorf("unable to get last insert id: %v", err)
	}
	return models.User{
		ID:        string(id),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

func (db *DB) UpdateUser(id string, user *models.User) (models.User, error) {
	err := db.QueryRow("UPDATE users SET first_name=$1, last_name=$2, email=$3 WHERE id = $4 RETURNING id, first_name, last_name, email", &user.FirstName, &user.LastName, &user.Email, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return models.User{}, fmt.Errorf("error updating user by ID: %v", err)
	}
	return *user, nil
}
