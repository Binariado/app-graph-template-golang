package repository

import (
	"app-graph/internal/domain/graph/model"
	"database/sql"
)

// Implementation of repositories to process data in a database

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db: db}
}

func (r *SQLUserRepository) FindByID(userID string) (*model.User, error) {
	if userID == "1" {
		var nameValue = "Name"
		return &model.User{
			ID:   "1",
			Name: nameValue,
		}, nil
	} else {
		return nil, nil
	}
}
