package usecase

import (
	"svc-users-go/module/v1/role/model"
	"svc-users-go/module/v1/role/repository"
	"svc-users-go/utils"
)

type roleUseCaseImpl struct {
	roleRepository repository.Repository
}

func NewRoleUseCase(roleRepos repository.Repository) UseCase {
	return &roleUseCaseImpl{
		roleRepository: roleRepos,
	}
}

func (ru *roleUseCaseImpl) CountAllRoles() (int64, error) {
	countAllRoleData, errorHandlerRepo := ru.roleRepository.Count()

	if !utils.GlobalErrorException(errorHandlerRepo) {
		return 0, errorHandlerRepo
	}

	return countAllRoleData, nil
}

func (ru *roleUseCaseImpl) FindAllRoles(limit int64, page int64) ([]model.Role, error) {
	var pages int64

	// Set paging per page
	if page == 1 {
		pages = page
	} else {
		pages = page * 10
	}

	findAllRoleRepo, errorHandlerRepo := ru.roleRepository.FindAll(limit, pages)
	if !utils.GlobalErrorException(errorHandlerRepo) {
		return nil, errorHandlerRepo
	}

	return findAllRoleRepo, nil
}

func (ru *roleUseCaseImpl) CreateNewRole(payload *model.CreateRole) error {
	errorHandlerSaveRepo := ru.roleRepository.Save(payload)
	if !utils.GlobalErrorDatabaseException(errorHandlerSaveRepo) {
		return errorHandlerSaveRepo
	}

	return nil
}
