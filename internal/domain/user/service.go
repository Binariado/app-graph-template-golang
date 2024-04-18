package user

import (
	"app-graph/internal/domain/graph/model"
	"context"
)

// Service handles business logic related to users.
type Service struct {
	repo Repository // Dependency towards the repository interface
}

// NewUserService creates a new instance of UserService.
func NewUserService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// FindByID search user for by id.
// It is recommended to configure the query functions in application/query (Left here only as an example of syntax for services)
// In services it is recommended to use it for data modification functions
func (s *Service) FindByID(ctx context.Context, userID string) (*model.User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
