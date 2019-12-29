package usecase

import "svc-users-go/module/v1/role/model"

type UseCase interface {
	CountAllRoles() (int64, error)
	FindAllRoles(limit int64, page int64) ([]model.Role, error)
	FindRoleById(id string) (model.Role, error)
	CreateNewRole(payload *model.CreateRole) error
}
