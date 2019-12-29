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

func (ru *roleUseCaseImpl) FindRoleById(id string) (model.Role, error) {
	findRoleById, errorHandlerRepo := ru.roleRepository.FindById(id)

	if !utils.GlobalErrorException(errorHandlerRepo) {
		return model.Role{}, errorHandlerRepo
	}

	return findRoleById, nil
}

func (ru *roleUseCaseImpl) UpdateRole(id string, payload *model.UpdateRole) error {
	// Check role by id exist or not
	_, errorHandlerRepo := ru.roleRepository.FindById(id)

	if !utils.GlobalErrorException(errorHandlerRepo) {
		return errorHandlerRepo
	}

	// Updating role user
	errorHandlerUpdateRole := ru.roleRepository.Update(id, payload)
	if !utils.GlobalErrorDatabaseException(errorHandlerUpdateRole) {
		return errorHandlerUpdateRole
	}

	return nil
}

func (ru *roleUseCaseImpl) DeleteRole(id string) error {
	// Check role by id exist or not
	_, errorHandlerRepo := ru.roleRepository.FindById(id)

	if !utils.GlobalErrorException(errorHandlerRepo) {
		return errorHandlerRepo
	}

	// Deleting role user
	errorHandlerDeleteRoleRepo := ru.roleRepository.Delete(id)
	if !utils.GlobalErrorException(errorHandlerDeleteRoleRepo) {
		return errorHandlerDeleteRoleRepo
	}

	return nil
}
