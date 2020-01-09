package usecase

import "svc-users-go/module/v1/access-control/model"

type UseCase interface {
	FindAllUser(limit int64, page int64) ([]model.AccessControl, error)
	CountAllUser() (int64, error)
	FindAccessControlById(id string) (model.DetailAccessControl, error)
	CreateNewAccessControl(payload *model.CreateAccessControl) (*model.CreateAccessControl, error)
}
