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

func (uc *userUseCaseImpl) FindAllUsers() ([]model.Users, error) {
	findAllUser, errorHandlerRepo := uc.userRepository.FindAll()

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
