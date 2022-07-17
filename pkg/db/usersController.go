package db

import (
	"context"
	"fmt"
	"net/mail"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string
	Password string
	Email    string
}

func (u User) ValidateEmail() bool {
	_, err := mail.ParseAddress(u.Email)
	return err == nil
}

func (u User) Insert() {
	if Client == nil {
		Init()
	}

	if !u.ValidateEmail() {
		panic("Invalid email")
	}
	hashed := "fakeHash" + u.Password
	u.Password = string(hashed)
	collection := Client.Database("test_database").Collection("users")
	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		panic("error during insert:" + err.Error())
	}

}

// func Insert(u *User) {
// 	if Client == nil {
// 		Init()
// 	}
// 	if !u.ValidateEmail() {
// 		panic("Invalid email")
// 	}
// 	fmt.Println("Inserting user:", u)

// 	collection := Client.Database("test_database").Collection("users")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	res, err := collection.InsertOne(ctx, u)
// 	if err != nil {
// 		panic("error during insert:" + err.Error())
// 	}
// 	fmt.Println("Inserted user with ID:", res.InsertedID)
// }

func (u User) Authenticate() {

	if Client == nil {
		Init()
	}

	var result User
	collection := Client.Database("test_database").Collection("users")
	hashed := "fakeHash" + u.Password
	u.Password = string(hashed)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"username": u.Username, "password": u.Password}).Decode(&result)
	if err != nil {
		panic("error during find:" + err.Error())
	}
	fmt.Println("Found user:", result)
}
