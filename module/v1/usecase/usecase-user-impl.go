package usecase

import (
	"svc-users-go/module/v1/repository"

	"svc-users-go/module/v1/model"

	"svc-users-go/utils"
)

type userUseCaseImpl struct {
	userRepository repository.Repository
}

func NewUserUseCase(UserRepos repository.Repository) UseCase {
	return &userUseCaseImpl{
		userRepository: UserRepos,
	}
}

func (uc *userUseCaseImpl) CountAllUsers() (int64, error) {
	countAllUserData, errorHandlerRepo := uc.userRepository.Count()

	if !utils.GlobalErrorException(errorHandlerRepo) {
		return 0, errorHandlerRepo
	}

	return countAllUserData, nil
}

func (uc *userUseCaseImpl) FindAllUsers(limit int64, page int64) ([]model.Users, error) {
	var pages int64

	// Set paging per page
	if page == 1 {
		pages = page
	} else {
		pages = page * 10
	}

	findAllUser, errorHandlerRepo := uc.userRepository.FindAll(limit, pages)

	if !utils.GlobalErrorDatabaseException(errorHandlerRepo) {
		return nil, errorHandlerRepo
	}

	return findAllUser, nil
}

func (uc *userUseCaseImpl) FindUserById(id string) (model.Users, error) {
	findUserById, errorHandlerRepo := uc.userRepository.FindById(id)

	if !utils.GlobalErrorDatabaseException(errorHandlerRepo) {
		return model.Users{}, errorHandlerRepo
	}

	return findUserById, nil

}

func (uc *userUseCaseImpl) CreateNewUser(payload *model.CreateUser) (*model.CreateUser, error) {

	errorHandlerRepo := uc.userRepository.Save(payload)

	if !utils.GlobalErrorDatabaseException(errorHandlerRepo) {
		return nil, errorHandlerRepo
	}

	return payload, nil
}

func (uc *userUseCaseImpl) UpdateUser(id string, payload *model.UpdateUser) error {

	// Find user first
	_, errorHandlerRepos := uc.userRepository.FindById(id)

	if !utils.GlobalErrorDatabaseException(errorHandlerRepos) {
		return errorHandlerRepos
	}

	// Update users
	errorHandlerRepo := uc.userRepository.Update(id, payload)

	if !utils.GlobalErrorDatabaseException(errorHandlerRepo) {
		return errorHandlerRepo
	}

	return nil
}

func (uc *userUseCaseImpl) DeleteUser(id string) error {
	// Find user first
	_, errorHandlerRepos := uc.userRepository.FindById(id)

	if !utils.GlobalErrorDatabaseException(errorHandlerRepos) {
		return errorHandlerRepos
	}

	errorHandlerRepo := uc.userRepository.Delete(id)

	if !utils.GlobalErrorDatabaseException(errorHandlerRepo) {
		return errorHandlerRepo
	}

	return nil
}
