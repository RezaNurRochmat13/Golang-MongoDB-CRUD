package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Age     int                `json:"age" bson:"age"`
	Address string             `json:"address" bson:"address"`
}

type CreateUser struct {
	Name    string `json:"name" bson:"name"`
	Age     int    `json:"age" bson:"age"`
	Address string `json:"address" bson:"address"`
}
type UpdateUser struct {
	Name    string `json:"name" bson:"name"`
	Age     int    `json:"age" bson:"age"`
	Address string `json:"address" bson:"address"`
}
