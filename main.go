package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_URL = "mongodb://localhost:27017"

func main() {
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      "dzikrurrohmani",
		Password:      "password",
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI(MONGODB_URL).SetAuth(credential)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connect, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected..")
	}

	defer func() {
		if err := connect.Disconnect(context.Background()); err != nil {
			panic(err)
		} else {
			fmt.Println("Disonnected..")
		}
	}()

	// membuat sebuah db - collection
	db := connect.Database("enigmago")
	coll := db.Collection("student")

	//Create
	newId, err := coll.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Jack"},
		{Key: "age", Value: 22},
		{Key: "gender", Value: "M"},
		{Key: "senior", Value: false},
	})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("inserted document with id: ", newId.InsertedID)
	students := []interface{}{
		bson.D{
			{Key: "name", Value: "Jack"},
			{Key: "age", Value: 23},
			{Key: "gender", Value: "M"},
			{Key: "senior", Value: false},
		},
		bson.D{
			{Key: "name", Value: "Jack"},
			{Key: "age", Value: 24},
			{Key: "gender", Value: "M"},
			{Key: "senior", Value: false},
		},
		bson.D{
			{Key: "name", Value: "Jack"},
			{Key: "age", Value: 25},
			{Key: "gender", Value: "M"},
			{Key: "senior", Value: false},
		},
	}
	out, err := coll.InsertMany(ctx, students)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("inserted document with id: ", out.InsertedIDs)
}

// Buat koneksi ke mongodd (url) -> mongodb://localhost:27017
// siapkan user auth -> username dan password
