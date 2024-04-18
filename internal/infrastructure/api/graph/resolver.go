package graph

import "app-graph/internal/domain/user"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserApp *user.Service
}
