package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string
	Price    int64
	Category string
}
