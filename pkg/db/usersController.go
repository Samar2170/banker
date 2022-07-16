package db

import (
	"context"
	"net/mail"
)

type User struct {
	username string
	password string
	email    string
}

func (u User) validate_email() bool {
	_, err := mail.ParseAddress(u.email)
	return err == nil
}

func (u User) Insert() {
	if !u.validate_email() {
		panic("Invalid email")
	}
	collection := DB.Collection("users")
	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		panic(err)
	}
}
