package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	coremodels "github.com/vivekprm/go-corelib/models"
	"github.com/vivekprm/go-microservices-full-stack/user-service/db"
	"github.com/vivekprm/go-microservices-full-stack/user-service/models"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	db *db.DB
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	listUserRe   = regexp.MustCompile(`^\/api\/users[\/]*$`)
	getUserRe    = regexp.MustCompile(`^\/api\/users\/(\d+)$`)
	createUserRe = regexp.MustCompile(`^\/api\/users[\/]*$`)
	updateUserRe = regexp.MustCompile(`^\/api\/users\/(\d+)$`)
	loginRe      = regexp.MustCompile(`^\/api\/login[\/]*$`)
)

func (uh *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	switch {
	case r.Method == http.MethodGet && listUserRe.MatchString(r.URL.Path):
		uh.List(w, r)
	case r.Method == http.MethodGet && getUserRe.MatchString(r.URL.Path):
		uh.Get(w, r)
	case r.Method == http.MethodPost && createUserRe.MatchString(r.URL.Path):
		uh.Create(w, r)
	case r.Method == http.MethodPut && updateUserRe.MatchString(r.URL.Path):
		uh.Update(w, r)
	case r.Method == http.MethodPost && loginRe.MatchString(r.URL.Path):
		uh.Login(w, r)
	default:
		notFound(w, r)
		return
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}
	var loginData loginRequest
	err = json.Unmarshal(body, &loginData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}
	user, err := uh.db.GetUserByEmail(loginData.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user details from email: %v", err), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, fmt.Sprintf("Error getting user details from email: %v", err), http.StatusNotFound)
		return
	}
	// Hashing the password with the default cost of 10
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid credentails: %v", err), http.StatusUnauthorized)
		return
	}
	log.Println("login successful.")
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	tk := &coremodels.Token{
		UserID: user.ID,
		Name:   fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		http.Error(w, fmt.Sprintf("Error creating jwt token: %v", err), http.StatusInternalServerError)
		return
	}
	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	json.NewEncoder(w).Encode(resp)
}

func (uh *userHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := uh.db.GetUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user list: %v", err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (uh *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	parts := getUserRe.FindStringSubmatch(r.URL.Path) // ["/api/users/1", "1"]
	if len(parts) <= 1 {
		http.Error(w, "Error getting id", http.StatusInternalServerError)
	}
	user, err := uh.db.GetUserByID(parts[1])
	if err != nil {
		http.Error(w, fmt.Sprintf("Error in getting user by id: %v", err), http.StatusInternalServerError)
	}
	if user == nil {
		http.Error(w, "User doesn't exist", http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(user)
}

func (uh *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}
	newUser, err := uh.db.AddUser(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(newUser)
}

func (uh *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}
	parts := getUserRe.FindStringSubmatch(r.URL.Path) // ["/api/users/1", "1"]
	if len(parts) <= 1 {
		http.Error(w, "Error getting id", http.StatusInternalServerError)
	}
	u, err := uh.db.UpdateUser(parts[1], &user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error in updating user: %v", err), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(u)
}
