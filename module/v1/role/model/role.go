package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	RoleID   primitive.ObjectID `json:"_id" bson:"_id"`
	RoleName string             `json:"role_name" bson:"role"`
}

type CreateRole struct {
	RoleName string `json:"role_name" bson:"role"`
}
