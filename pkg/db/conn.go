package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://samar2:tests_1234@127.0.0.1:27017/test_database?authSource=admin&readPreference=primary&directConnection=true&ssl=false"

var Client *mongo.Client

func Init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic("Error connecting to MongoDB: " + err.Error())
	}
	Client = client
	// DB = client.Database("test_database")
}

// func CheckDBConnection() {
// 	if DB == nil {
// 		Init()
// 	}
// 	if err := Client.Ping(context.TODO(), nil); err != nil {
// 		panic("Error pinging MongoDB: " + err.Error())
// 	}

// }
