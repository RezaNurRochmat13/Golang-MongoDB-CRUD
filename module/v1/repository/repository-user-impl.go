package repository

import (
	"context"
	"log"
	"svc-users-go/module/v1/model"

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

	if errorCsr != nil {
		log.Println("Error query mongo : ", errorCsr.Error())
		return nil, errorCsr
	}

	for csr.Next(cntx) {
		errorHandlerDecodeData := csr.Decode(&user)

		if errorHandlerDecodeData != nil {
			log.Println("Error decode data : ", errorHandlerDecodeData.Error())
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

	if errorGetOneUser != nil {
		log.Println("Error when get one : ", errorGetOneUser.Error())
		return model.Users{}, errorGetOneUser
	}

	return user, nil

}

func (ur *userRepositoryImpl) Save(payload *model.CreateUser) error {
	collection := ur.Connection.Database("authentication")

	_, errorHandlerSaveUser := collection.Collection("users").InsertOne(cntx, payload)

	if errorHandlerSaveUser != nil {
		log.Println("Error when saving mongo : ", errorHandlerSaveUser.Error())
		return errorHandlerSaveUser
	}

	return nil
}