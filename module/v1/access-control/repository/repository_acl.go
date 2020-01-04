package repository

import (
	"context"
	"svc-users-go/module/v1/access-control/model"
	"svc-users-go/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cntx context.Context

type accessControlRepositoryImpl struct {
	Connection *mongo.Database
}

func NewAccessControlRepository(Connection *mongo.Database) Repository {
	return &accessControlRepositoryImpl{Connection: Connection}
}

func (ar *accessControlRepositoryImpl) FindAll(limit int64, page int64) ([]model.AccessControl, error) {
	var (
		aclModel       model.AccessControl
		aclModelResult []model.AccessControl
		csr            *mongo.Cursor
		filterOptions  = options.Find()
		errorCsr       error
	)

	filterOptions.SetSkip(page)
	filterOptions.SetLimit(limit)

	csr, errorCsr = ar.Connection.Collection("access_control").Find(cntx, bson.M{}, filterOptions)
	if !utils.GlobalErrorDatabaseException(errorCsr) {
		return nil, errorCsr
	}

	for csr.Next(cntx) {
		errorHandlerDecode := csr.Decode(&aclModel)
		if !utils.GlobalErrorDatabaseException(errorHandlerDecode) {
			return nil, errorHandlerDecode
		}

		aclModelResult = append(aclModelResult, aclModel)
	}

	return aclModelResult, nil
}

func (ar *accessControlRepositoryImpl) Count() (int64, error) {
	countAllAccessControl, errorHandlerAcl := ar.Connection.Collection("access_control").CountDocuments(cntx, bson.M{})
	if !utils.GlobalErrorDatabaseException(errorHandlerAcl) {
		return 0, errorHandlerAcl
	}

	return countAllAccessControl, nil
}
