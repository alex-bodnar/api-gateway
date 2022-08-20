package repository

import (
	"context"

	"authenticator/internal/api/domain/users"
)

//go:generate mockgen -source repository.go -destination ./repository_mock.go -package repository

type (
	// Users - describe an interface for working with users.
	Users interface {
		GetByName(ctx context.Context, name string) (users.User, error)
		Save(ctx context.Context, userData users.User) error
	}
)
