package main

import "go-mongod/delivery"

func main() {
	delivery.NewServer().Run()
}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const MONGODB_URL = "mongodb://localhost:27017"

// type Student struct {
// 	// kalo gak pake struct tag, nanti jadi lowercase``
// 	Id       primitive.ObjectID `bson:"_id"`
// 	Name     string             `bson:"fullName"`
// 	Age      int                `bson:"age"`
// 	Gender   string             `bson:"gender"`
// 	JoinDate primitive.DateTime `bson:"joinDate"`
// 	Senior   bool               `bson:"senior"`
// }

// func main() {

// 	credential := options.Credential{
// 		AuthMechanism: "SCRAM-SHA-256",
// 		Username:      "dzikrurrohmani",
// 		Password:      "password",
// 	}
// 	clientOptions := options.Client()
// 	clientOptions.ApplyURI(MONGODB_URL).SetAuth(credential)
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	connect, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("Connected...")
// 	}
// 	defer func() {
// 		if err := connect.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	// membuat sebuah db - collection
// 	db := connect.Database("enigmago")
// 	coll := db.Collection("student")

// 	// Create
// 	// newId, err := coll.InsertOne(ctx, bson.D{
// 	// 	{"age", 19},
// 	// 	{"name", "Bulan"},
// 	// 	{"gender", "F"},
// 	// 	{"senior", false},
// 	// })

// 	// slice
// 	// jd01 := parseTime("2022-07-02 15:04:05")
// 	// jd02 := parseTime("2022-07-03 15:04:05")
// 	// // jd03 := parseTime("2022-07-04 15:04:05")
// 	// students := []interface{}{
// 	// 	bson.D{
// 	// 		{"name", "Sita"},
// 	// 		{"age", 29},
// 	// 		{"gender", "F"},
// 	// 		{"joinDate", primitive.NewDateTimeFromTime(jd01)},
// 	// 		{"senior", true},
// 	// 	},
// 	// 	bson.D{
// 	// 		{"name", "Melani"},
// 	// 		{"age", 25},
// 	// 		{"gender", "F"},
// 	// 		{"joinDate", jd02},
// 	// 		{"senior", true},
// 	// 	},
// 	// 	bson.D{
// 	// 		{"name", "Suci"},
// 	// 		{"age", 10},
// 	// 		{"gender", "F"},
// 	// 		{"joinDate", primitive.NewDateTimeFromTime(parseTime("2022-07-13"))},
// 	// 		{"senior", false},
// 	// 	},
// 	// }

// 	// // INSERT
// 	// newStudent := Student{
// 	// 	Id:       primitive.NewObjectID(),
// 	// 	Name:     "Doni",
// 	// 	Age:      23,
// 	// 	Gender:   "M",
// 	// 	JoinDate: primitive.NewDateTimeFromTime(parseTime("2022-07-13 00:00:00")),
// 	// 	Senior:   false,
// 	// }

// 	// newId, err := coll.InsertOne(ctx, newStudent)

// 	// if err != nil {
// 	// 	log.Println(err.Error())
// 	// }
// 	// fmt.Printf("inserted document with ID %v\n", newId.InsertedID)
// 	// // UPDATE
// 	// idUpdate, _ := primitive.ObjectIDFromHex("62cd19e4084f53efca1a2502")
// 	// filter := bson.D{{"_id", idUpdate}}
// 	// update := bson.D{{"$set", bson.D{{"fullName", "Dzikrur"}}}}
// 	// _, err = coll.UpdateOne(ctx, filter, update)
// 	// if err != nil {
// 	// 	log.Println(err.Error())
// 	// }
// 	// // DELETE
// 	// idDelete, _ := primitive.ObjectIDFromHex("62cd19e4084f53efca1a2502")
// 	// filter := bson.D{{"_id", idDelete}}
// 	// _, err = coll.DeleteOne(ctx, filter)
// 	// if err != nil {
// 	// 	log.Println(err.Error())
// 	// }
// 	// // READ
// 	// // select * from student
// 	// cursor, err := coll.Find(ctx, bson.D{})
// 	// if err != nil {
// 	// 	log.Println(err.Error())
// 	// }
// 	// var students []bson.D
// 	// err = cursor.All(ctx, &students)
// 	// if err != nil {
// 	// 	log.Println(err.Error())
// 	// }
// 	// for _, student := range students {
// 	// 	fmt.Println(student)
// 	// }

// 	// // PROJECTION
// 	// opts := options.Find().SetProjection(bson.D{
// 	// 	{Key: "_id", Value: 0},
// 	// 	{Key: "fullName", Value: 1},
// 	// 	{Key: "age", Value: 1},
// 	// })
// 	// cursor, err := coll.Find(ctx, bson.D{}, opts)
// 	// if err != nil {
// 	// 	log.Println(err.Error())
// 	// }
// 	// var students []bson.D
// 	// err = cursor.All(ctx, &students)
// 	// if err != nil {
// 	// 	log.Println(err.Error())
// 	// }
// 	// for _, student := range students {
// 	// 	fmt.Println(student)
// 	// }

// 	// // Filter genderAndAGe
// 	// filtered := bson.D{
// 	// 	{Key: "$and", Value: bson.A{
// 	// 		bson.D{
// 	// 			{Key: "gender", Value: "M"},
// 	// 			{Key: "age", Value: bson.D{{Key: "$gte", Value: 24}}},
// 	// 		},
// 	// 	}},
// 	// }
// 	// projection := options.Find().SetProjection(bson.D{
// 	// 	{Key: "_id", Value: 0},
// 	// 	{Key: "fullName", Value: 1},
// 	// 	{Key: "age", Value: 1},
// 	// })
// 	// cursor, err := coll.Find(ctx, filtered, projection)
// 	// if err != nil {
// 	// 	log.Println(err.Error())
// 	// }
// 	// var students []bson.D
// 	// err = cursor.All(ctx, &students)
// 	// if err != nil {
// 	// 	log.Println(err.Error())
// 	// }
// 	// for _, student := range students {
// 	// 	fmt.Println(student)
// 	// }

// 	// Maping to struct
// 	filterGenderAndAge := make([]*Student, 0)
// 	filtered := bson.D{
// 		{Key: "$and", Value: bson.A{
// 			bson.D{
// 				{Key: "gender", Value: "M"},
// 				{Key: "age", Value: bson.D{{Key: "$gte", Value: 24}}},
// 			},
// 		}},
// 	}
// 	projection := options.Find().SetProjection(bson.D{
// 		{Key: "_id", Value: 1},
// 		{Key: "fullName", Value: 1},
// 		{Key: "gender", Value: 1},
// 		{Key: "age", Value: 1},
// 	})
// 	cursor, err := coll.Find(ctx, filtered, projection)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	for cursor.Next(ctx) {
// 		var student Student
// 		err := cursor.Decode(&student)
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		filterGenderAndAge = append(filterGenderAndAge, &student)
// 	}
// 	for _, student := range filterGenderAndAge {
// 		fmt.Println(student)
// 	}

// 	// Aggregation
// 	coll2 := connect.Database("enigma").Collection("products")
// 	count, err := coll2.CountDocuments(ctx, bson.D{})
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	fmt.Println("all", count)

// 	// filter
// 	count, err = coll2.CountDocuments(ctx, bson.D{{Key: "category", Value: "food"}})
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	fmt.Println("food", count)

// 	// match, group, sort, dll
// 	matchStage := bson.D{
// 		{Key: "$match", Value: bson.D{
// 			{Key: "category", Value: "food"},
// 		}},
// 	}

// 	groupStage := bson.D{
// 		{Key: "$group", Value: bson.D{
// 			{Key: "_id", Value: "$category"},
// 			{Key: "Total", Value: bson.D{{Key: "$sum", Value: 1}}},
// 		}},
// 	}

// 	cursor, err = coll2.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	var productCount []bson.M
// 	err = cursor.All(ctx, &productCount)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	for _, product := range productCount {
// 		fmt.Printf("Group[%v], Total[%v]\n",product["_id"], product["Total"])
// 	}
// }

// func parseTime(date string) time.Time {
// 	layoutFormat := "2006-01-02 15:04:05"
// 	parse, _ := time.Parse(layoutFormat, date)
// 	return parse
// }

// /**
// mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb
// * Buat koneksi ke mongodb (url) -> mongodb://localhost:27017
// * Siapkan User Auth: username dan password
// */
