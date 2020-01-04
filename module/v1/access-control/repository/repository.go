package repository

import "svc-users-go/module/v1/access-control/model"

type Repository interface {
	FindAll(limit int64, page int64) ([]model.AccessControl, error)
	Count() (int64, error)
}
