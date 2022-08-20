package repository

import (
	"context"

	"user/internal/api/domain/users"
)

//go:generate mockgen -source repository.go -destination ./repository_mock.go -package repository

type (
	// Authenticator - describe an interface for working with authenticator service.
	Authenticator interface {
		SendNewUser(ctx context.Context, user users.User) error
	}

	// Users - describe an interface for working with users database.
	Users interface {
		GetByName(ctx context.Context, name string) (users.User, error)
		Create(ctx context.Context, userData users.User) (users.User, error)
	}
)
