package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Aloha World!"))
}

func main() {

	index := http.HandlerFunc(hello)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", index)
	r.Post("/signup", signup)
	http.ListenAndServe(":8080", r)

}

func signup(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	if username == "" || password == "" || email == "" {
		w.Write([]byte("Missing username, password or email"))
		return
	}
	user := User{username, password, email}
	if !user.validate_email() {
		w.Write([]byte("Invalid email"))
		return
	}
	w.Write([]byte("User created"))

}
