package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/vivekprm/go-microservices-full-stack/todo-service/config"
	"github.com/vivekprm/go-microservices-full-stack/todo-service/models"
)

type Store struct {
	*sql.DB
}

func ConnectDB(cfg *config.Config) (*Store, error) {
	db, err := sql.Open(cfg.DriverName, cfg.GetDBConnectionString())
	if err != nil {
		log.Fatalf("unable to open connection to database: %v", err)
		return nil, fmt.Errorf("unable to open connection to database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	log.Println("Database connected.")
	return &Store{db}, nil
}

func (db *Store) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	var id int
	var ts time.Time
	err := db.QueryRow("insert into todos (name, description, status, created_by) VALUES ($1, $2, $3, $4) RETURNING id, created_on", todo.Name, todo.Description, todo.Status, todo.CreatedBy).Scan(&id, &ts)
	if err != nil {
		return &models.Todo{}, fmt.Errorf("error in creating todo: %v", err)
	}
	return &models.Todo{
		ID:          string(id),
		Name:        todo.Name,
		Description: todo.Description,
		Status:      todo.Status,
		CreatedBy:   todo.CreatedBy,
		CreatedOn:   ts.String(),
	}, nil
}

func (db *Store) ListTodo() ([]models.Todo, error) {
	rows, err := db.Query("select * from todos")
	if err != nil {
		return nil, fmt.Errorf("error in getting todos %v", err)
	}
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.Status, &todo.CreatedBy, &todo.CreatedOn)
		if err != nil {
			return nil, fmt.Errorf("error in scanning the row %v", err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
