package usecase

import "svc-users-go/module/v1/model"

type UseCase interface {
	FindAllUsers() ([]model.Users, error)
	FindUserById(id string) (model.Users, error)
	CreateNewUser(payload *model.CreateUser) (*model.CreateUser, error)
}