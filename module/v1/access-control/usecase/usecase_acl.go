package usecase

import (
	"svc-users-go/module/v1/access-control/model"
	"svc-users-go/module/v1/access-control/repository"
	"svc-users-go/utils"
)

type accessControlUseCaseImpl struct {
	accessControlRepository repository.Repository
}

func NewAccessControlUseCase(accessControlRepo repository.Repository) UseCase {
	return &accessControlUseCaseImpl{accessControlRepository: accessControlRepo}
}

func (au *accessControlUseCaseImpl) FindAllUser(limit int64, page int64) ([]model.AccessControl, error) {
	var pages int64

	// Set paging per page
	if page == 1 {
		pages = page
	} else {
		pages = page * 10
	}

	findAllAccessControl, errorHandlerRepo := au.accessControlRepository.FindAll(limit, pages)
	if !utils.GlobalErrorDatabaseException(errorHandlerRepo) {
		return nil, errorHandlerRepo
	}

	return findAllAccessControl, nil
}

func (au *accessControlUseCaseImpl) CountAllUser() (int64, error) {
	countAllUserRepo, errorHandlerRepo := au.accessControlRepository.Count()
	if !utils.GlobalErrorDatabaseException(errorHandlerRepo) {
		return 0, errorHandlerRepo
	}

	return countAllUserRepo, nil
}

func (au *accessControlUseCaseImpl) FindAccessControlById(id string) (model.DetailAccessControl, error) {
	findAccessControlById, errorHandlerRepo := au.accessControlRepository.FindById(id)
	if !utils.GlobalErrorException(errorHandlerRepo) {
		return model.DetailAccessControl{}, errorHandlerRepo
	}

	return findAccessControlById, errorHandlerRepo
}

func (au *accessControlUseCaseImpl) CreateNewAccessControl(payload *model.CreateAccessControl) (*model.CreateAccessControl, error) {
	_, errorHandlerRepo := au.accessControlRepository.Save(payload)
	if !utils.GlobalErrorException(errorHandlerRepo) {
		return nil, errorHandlerRepo
	}

	return payload, nil
}
