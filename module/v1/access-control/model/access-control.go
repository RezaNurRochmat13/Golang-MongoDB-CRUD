package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccessControl struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	AccessControlName string             `json:"access_control_name" bson:"access_control_name"`
}

type DetailAccessControl struct {
	Id                       primitive.ObjectID `json:"id" bson:"_id"`
	AccessControlName        string             `json:"access_control_name" bson:"access_control_name"`
	AccessControlDescription string             `json:"access_control_description" bson:"access_control_description"`
}
