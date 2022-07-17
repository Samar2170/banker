package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samar2170/banker/pkg/db"
)

var CollectionError = errors.New("Collection not found")
var CusorError = errors.New("Cursor not found")

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))

}

func checkDB() {
	if db.Client == nil {
		db.Init()
	}
	if err := db.Client.Ping(context.TODO(), nil); err != nil {
		panic("Error pinging MongoDB: " + err.Error())
	}
}

func main() {
	index := http.HandlerFunc(hello)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", index)
	r.Post("/signup", signup)
	r.Post("/login", login)
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
	user := db.User{Username: username, Password: password, Email: email}
	// db.Insert(&user)
	user.Insert()
	w.Write([]byte("User created"))

}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		w.Write([]byte("Missing username or password"))
		return
	}
	user := db.User{Username: username, Password: password}
	// db.Insert(&user)
	user.Authenticate()
	w.Write([]byte("Logged in"))

}
