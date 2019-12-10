package repository

import "svc-users-go/module/v1/model"

type Repository interface {
	FindAll() ([]model.Users, error)
	FindById(id string) (model.Users, error)
	Save(payload *model.CreateUser) error
	Update(id string, payload *model.UpdateUser) error
	Delete(id string) error
}
