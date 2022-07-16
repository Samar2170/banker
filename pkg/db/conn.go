package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://samar:tests_123@127.0.0.1:27017/test?authSource=admin&readPreference=primary&ssl=false"

var Client *mongo.Client

var DB *mongo.Database

func Init() {
	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	DB = Client.Database("test")

}
func CheckDBConnection() {
	if DB == nil {
		Init()
	}
	if err := Client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}

}
