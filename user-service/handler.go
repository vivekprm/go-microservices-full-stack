package main

import (
	"net/http"
	"regexp"
)

type userHandler struct {
}

var (
	listUserRe   = regexp.MustCompile(`^\/api\/users[\/]*$`)
	getUserRe    = regexp.MustCompile(`^\/api\/users\/(\d+)$`)
	createUserRe = regexp.MustCompile(`^\/api\/users[\/]*$`)
	updateUserRe = regexp.MustCompile(`^\/api\/users\/(\d+)$`)
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
	}
}

func (uh *userHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List API"))
}

func (uh *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get API"))
}

func (uh *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create API"))
}

func (uh *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update API"))
}
