package repository

import "svc-users-go/module/v1/role/model"

type Repository interface {
	Count() (int64, error)
	FindAll(limit int64, page int64) ([]model.Role, error)
	FindById(id string) (model.Role, error)
	Save(rolePayload *model.CreateRole) error
}
