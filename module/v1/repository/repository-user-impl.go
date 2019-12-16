package repository

import (
	"context"
	"svc-users-go/module/v1/model"
	"svc-users-go/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cntx context.Context

type userRepositoryImpl struct {
	Connection *mongo.Client
}

func NewUserRepository(Connection *mongo.Client) Repository {
	return &userRepositoryImpl{Connection: Connection}
}

func (ur *userRepositoryImpl) Count() (int64, error) {
	collection := ur.Connection.Database("authentication")

	countRecord, errorHandlerCount := collection.Collection("users").CountDocuments(cntx, bson.M{}, nil)

	if !utils.GlobalErrorDatabaseException(errorHandlerCount) {
		return 0, errorHandlerCount
	}

	return countRecord, nil
}

func (ur *userRepositoryImpl) FindAll(limit int64, offset int64) ([]model.Users, error) {
	collection := ur.Connection.Database("authentication")

	var (
		user     model.Users
		users    []model.Users
		csr      *mongo.Cursor
		errorCsr error
	)

	filterOptions := options.Find()
	filterOptions.SetLimit(limit)
	filterOptions.SetSkip(offset)

	csr, errorCsr = collection.Collection("users").Find(cntx, bson.M{}, filterOptions)

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

func (ur *userRepositoryImpl) Delete(id string) error {
	var (
		database    = ur.Connection.Database("authentication")
		objectID, _ = primitive.ObjectIDFromHex(id)
	)

	filter := bson.M{"_id": objectID}

	_, errorHandlerDelete := database.Collection("users").DeleteOne(cntx, filter)

	if !utils.GlobalErrorDatabaseException(errorHandlerDelete) {
		return errorHandlerDelete
	}

	return nil
}
