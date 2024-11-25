package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	coremodels "github.com/vivekprm/go-corelib/models"
	"github.com/vivekprm/go-microservices-full-stack/todo-service/db"
	"github.com/vivekprm/go-microservices-full-stack/todo-service/models"
)

type TodoHandler struct {
	db *db.Store
}

var (
	listTodoRe   = regexp.MustCompile(`^\/api\/todos[\/]*$`)
	createTodoRe = regexp.MustCompile(`^\/api\/todos[\/]*$`)
	getTodoRe    = regexp.MustCompile(`^\/api\/todos\/(\d+)$`)
	updateTodoRe = regexp.MustCompile(`^\/api\/todos\/(\d+)$`)
	deleteTodoRe = regexp.MustCompile(`^\/api\/todos\/(\d+)$`)
)

func (th *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && listTodoRe.MatchString(r.URL.Path):
		th.listTodo(w, r)
	case r.Method == http.MethodPost && createTodoRe.MatchString(r.URL.Path):
		th.createTodo(w, r)
	case r.Method == http.MethodGet && getTodoRe.MatchString(r.URL.Path):
		th.getTodo(w, r)
	case r.Method == http.MethodPut && updateTodoRe.MatchString(r.URL.Path):
		th.updateTodo(w, r)
	case r.Method == http.MethodDelete && deleteTodoRe.MatchString(r.URL.Path):
		th.deleteTodo(w, r)
	default:
		notFound(w)
	}
}

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func (th *TodoHandler) listTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := th.db.ListTodo()
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to list todos: %v", err), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(todos)
}

func (th *TodoHandler) createTodo(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("user")
	fmt.Println(jwt)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error in reading request body: %v", err), http.StatusBadRequest)
	}
	var todo *models.Todo
	err = json.Unmarshal(data, &todo)
	if err != nil {
		http.Error(w, fmt.Sprintf("error in unmarshalling: %v", err), http.StatusInternalServerError)
	}
	todo.Status = models.Pending
	todo.CreatedBy = jwt.(*coremodels.Token).UserID
	resp, err := th.db.CreateTodo(todo)
	if err != nil {
		http.Error(w, fmt.Sprintf("error in creating todo: %v", err), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(resp)
}

func (th *TodoHandler) getTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting todo"))
}

func (th *TodoHandler) updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating todo"))
}

func (th *TodoHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting todo"))
}
