package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_URL = "mongodb://localhost:27017"

func main() {

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      "utsmanarifin",
		Password:      "P@ssw0rd",
	}
	clientOptions := options.Client()
	clientOptions.ApplyURI(MONGODB_URL).SetAuth(credential)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connect, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected...")
	}
	defer func() {
		if err := connect.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// membuat sebuah db - collection
	db := connect.Database("enigma")
	coll := db.Collection("student")

	// Create
	// newId, err := coll.InsertOne(ctx, bson.D{
	// 	{"age", 19},
	// 	{"name", "Bulan"},
	// 	{"gender", "F"},
	// 	{"senior", false},
	// })

	// slice
	jd01 := parseTime("2022-07-02 15:04:05")
	jd02 := parseTime("2022-07-03 15:04:05")
	jd03 := parseTime("2022-07-04 15:04:05")
	students := []interface{}{
		bson.D{
			{Key: "name", Value: "Sita"},
			{Key: "age", Value: 29},
			{Key: "gender", Value: "F"},
			{Key: "joinDate", Value: primitive.NewDateTimeFromTime(jd01)},
			{Key: "senior", Value: true},
		},
		bson.D{
			{Key: "name", Value: "Melani"},
			{Key: "age", Value: 25},
			{Key: "gender", Value: "F"},
			{Key: "joinDate", Value: jd02},
			{Key: "senior", Value: true},
		},
		bson.D{
			{Key: "name", Value: "Suci"},
			{Key: "age", Value: 10},
			{Key: "gender", Value: "F"},
			{Key: "joinDate", Value: jd03},
			{Key: "senior", Value: false},
		},
	}
	newId, err := coll.InsertMany(ctx, students)

	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("inserted document with ID %v\n", newId.InsertedIDs)

}

func parseTime(date string) time.Time {
	layoutFormat := "2006-01-02 15:04:05"
	parse, _ := time.Parse(layoutFormat, date)
	return parse
}

// Buat koneksi ke mongodd (url) -> mongodb://localhost:27017
// siapkan user auth -> username dan password
