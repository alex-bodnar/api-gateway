package services

import (
	"context"

	"authenticator/internal/api/domain/users"
)

//go:generate mockgen -source services.go -destination ./services_mock.go -package services

type (
	// UsersService - describe an interface for working with users.
	UsersService interface {
		CheckUserByName(ctx context.Context, name string) error
		SaveNewUser(ctx context.Context, userData users.User) error
	}
)
