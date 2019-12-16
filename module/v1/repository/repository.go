package repository

import "svc-users-go/module/v1/model"

type Repository interface {
	FindAll(limit int64, offset int64) ([]model.Users, error)
	Count() (int64, error)
	FindById(id string) (model.Users, error)
	Save(payload *model.CreateUser) error
	Update(id string, payload *model.UpdateUser) error
	Delete(id string) error
}
