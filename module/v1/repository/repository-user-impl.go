package repository

import (
	"context"
	"svc-users-go/module/v1/model"
	"svc-users-go/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cntx context.Context

type userRepositoryImpl struct {
	Connection *mongo.Client
}

func NewUserRepository(Connection *mongo.Client) Repository {
	return &userRepositoryImpl{Connection: Connection}
}

func (ur *userRepositoryImpl) FindAll() ([]model.Users, error) {
	collection := ur.Connection.Database("authentication")

	var (
		user  model.Users
		users []model.Users
	)

	csr, errorCsr := collection.Collection("users").Find(cntx, bson.D{{}}, nil)

	if !utils.GlobalErrorDatabaseException(errorCsr) {
		return nil, errorCsr
	}

	for csr.Next(cntx) {
		errorHandlerDecodeData := csr.Decode(&user)

		if !utils.GlobalErrorDatabaseException(errorHandlerDecodeData) {
			return nil, errorHandlerDecodeData
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *userRepositoryImpl) FindById(id string) (model.Users, error) {
	collection := ur.Connection.Database("authentication")
	userId, _ := primitive.ObjectIDFromHex(id)

	var (
		user   model.Users
		filter = bson.M{"_id": userId}
	)

	errorGetOneUser := collection.Collection("users").FindOne(cntx, filter).Decode(&user)

	if !utils.GlobalErrorDatabaseException(errorGetOneUser) {
		return model.Users{}, errorGetOneUser
	}

	return user, nil

}

func (ur *userRepositoryImpl) Save(payload *model.CreateUser) error {
	collection := ur.Connection.Database("authentication")

	_, errorHandlerSaveUser := collection.Collection("users").InsertOne(cntx, payload)

	if !utils.GlobalErrorDatabaseException(errorHandlerSaveUser) {
		return errorHandlerSaveUser
	}

	return nil
}

func (ur *userRepositoryImpl) Update(id string, payload *model.UpdateUser) error {
	collection := ur.Connection.Database("authentication")
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id": objectID,
	}

	updateField := bson.M{
		"$set": bson.M{
			"name":    payload.Name,
			"age":     payload.Age,
			"address": payload.Address,
		}}

	_, errorHandlerUpdateUser := collection.Collection("users").UpdateOne(cntx, filter, updateField)

	if !utils.GlobalErrorDatabaseException(errorHandlerUpdateUser) {
		return errorHandlerUpdateUser
	}

	return nil
}
