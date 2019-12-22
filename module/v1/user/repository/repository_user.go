package repository

import (
	"context"
	"svc-users-go/module/v1/user/model"
	"svc-users-go/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cntx context.Context

type userRepositoryImpl struct {
	Connection *mongo.Database
}

func NewUserRepository(Connection *mongo.Database) Repository {
	return &userRepositoryImpl{Connection: Connection}
}

func (ur *userRepositoryImpl) Count() (int64, error) {
	countRecord, errorHandlerCount := ur.Connection.Collection("users").CountDocuments(cntx, bson.M{}, nil)

	if !utils.GlobalErrorDatabaseException(errorHandlerCount) {
		return 0, errorHandlerCount
	}

	return countRecord, nil
}

func (ur *userRepositoryImpl) FindAll(name string, limit int64, offset int64) ([]model.Users, error) {
	var (
		user          model.Users
		users         []model.Users
		filterOptions = options.Find()
		csr           *mongo.Cursor
		errorCsr      error
	)

	filterOptions.SetLimit(limit)
	filterOptions.SetSkip(offset)

	if name != "" {
		csr, errorCsr = ur.Connection.Collection("users").Find(cntx, bson.M{"name": name}, filterOptions)
		if !utils.GlobalErrorDatabaseException(errorCsr) {
			return nil, errorCsr
		}
	} else {
		csr, errorCsr = ur.Connection.Collection("users").Find(cntx, bson.M{}, filterOptions)

		if !utils.GlobalErrorDatabaseException(errorCsr) {
			return nil, errorCsr
		}
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
	var (
		user      model.Users
		userId, _ = primitive.ObjectIDFromHex(id)
		filter    = bson.M{"_id": userId}
	)

	errorGetOneUser := ur.Connection.Collection("users").FindOne(cntx, filter).Decode(&user)

	if !utils.GlobalErrorDatabaseException(errorGetOneUser) {
		return model.Users{}, errorGetOneUser
	}

	return user, nil

}

func (ur *userRepositoryImpl) Save(payload *model.CreateUser) error {
	_, errorHandlerSaveUser := ur.Connection.Collection("users").InsertOne(cntx, payload)

	if !utils.GlobalErrorDatabaseException(errorHandlerSaveUser) {
		return errorHandlerSaveUser
	}

	return nil
}

func (ur *userRepositoryImpl) Update(id string, payload *model.UpdateUser) error {
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

	_, errorHandlerUpdateUser := ur.Connection.Collection("users").UpdateOne(cntx, filter, updateField)

	if !utils.GlobalErrorDatabaseException(errorHandlerUpdateUser) {
		return errorHandlerUpdateUser
	}

	return nil
}

func (ur *userRepositoryImpl) Delete(id string) error {
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectID}

	_, errorHandlerDelete := ur.Connection.Collection("users").DeleteOne(cntx, filter)

	if !utils.GlobalErrorDatabaseException(errorHandlerDelete) {
		return errorHandlerDelete
	}

	return nil
}
