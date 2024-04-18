package user

import "app-graph/internal/domain/graph/model"

type Repository interface {
	FindByID(id string) (*model.User, error)
}
