package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Books struct {
	Name   string             `bson: "name_books`
	Author string             `bson: "author_books`
	Price  int                `bson: "price"`
	ID     primitive.ObjectID `bson:"_id"`
}
