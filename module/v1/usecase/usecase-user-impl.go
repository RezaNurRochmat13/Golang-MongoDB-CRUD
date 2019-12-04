package usecase

import "svc-users-go/module/v1/repository"

import "svc-users-go/module/v1/model"

import "log"

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

	if errorHandlerRepo != nil {
		log.Println("Error when get repo : ", errorHandlerRepo.Error())
		return nil, errorHandlerRepo
	}

	return findAllUser, nil
}

func (uc *userUseCaseImpl) FindUserById(id string) (model.Users, error) {
	findUserById, errorHandlerRepo := uc.userRepository.FindById(id)

	if errorHandlerRepo != nil {
		log.Println("Error when get repo : ", errorHandlerRepo.Error())
		return model.Users{}, errorHandlerRepo
	}

	return findUserById, nil

}

func (uc *userUseCaseImpl) CreateNewUser(payload *model.CreateUser) (*model.CreateUser, error) {

	errorHandlerRepo := uc.userRepository.Save(payload)

	if errorHandlerRepo != nil {
		log.Println("Error when get repo : ", errorHandlerRepo.Error())
		return nil, errorHandlerRepo
	}

	return payload, nil
}

func (uc *userUseCaseImpl) UpdateUser(id string, payload *model.UpdateUser) error {

	// Find user first
	_, errorHandlerRepos := uc.userRepository.FindById(id)

	if errorHandlerRepos != nil {
		log.Println("User not found", errorHandlerRepos.Error())
		return errorHandlerRepos
	}

	// Update users
	errorHandlerRepo := uc.userRepository.Update(id, payload)

	if errorHandlerRepo != nil {
		log.Println("Error when get repo : ", errorHandlerRepo.Error())
		return errorHandlerRepo
	}

	return nil
}
