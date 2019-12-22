package repository

import (
	"context"
	"svc-users-go/module/v1/role/model"
	"svc-users-go/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cntx context.Context

type roleRepositoryImpl struct {
	Connection *mongo.Database
}

func NewRoleRepository(Connection *mongo.Database) Repository {
	return &roleRepositoryImpl{Connection: Connection}
}

func (ru *roleRepositoryImpl) Count() (int64, error) {
	countRecord, errorHandlerCount := ru.Connection.Collection("role").CountDocuments(cntx, bson.M{}, nil)

	if !utils.GlobalErrorDatabaseException(errorHandlerCount) {
		return 0, errorHandlerCount
	}

	return countRecord, nil
}

func (ru *roleRepositoryImpl) FindAll(limit int64,
	page int64) ([]model.Role, error) {

	var (
		role          model.Role
		roles         []model.Role
		csr           *mongo.Cursor
		filterOptions = options.Find()
		errorCsr      error
	)

	filterOptions.SetLimit(limit)
	filterOptions.SetSkip(page)

	csr, errorCsr = ru.Connection.Collection("role").Find(cntx, bson.M{}, filterOptions)

	if !utils.GlobalErrorDatabaseException(errorCsr) {
		return nil, errorCsr
	}

	for csr.Next(cntx) {
		errorDecode := csr.Decode(&role)

		if !utils.GlobalErrorDatabaseException(errorDecode) {
			return nil, errorDecode
		}

		roles = append(roles, role)
	}

	return roles, nil

}
