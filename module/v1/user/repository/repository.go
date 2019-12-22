package repository

import "svc-users-go/module/v1/user/model"

type Repository interface {
	FindAll(name string, limit int64, offset int64) ([]model.Users, error)
	Count() (int64, error)
	FindById(id string) (model.Users, error)
	Save(payload *model.CreateUser) error
	Update(id string, payload *model.UpdateUser) error
	Delete(id string) error
}
