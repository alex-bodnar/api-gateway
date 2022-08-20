package services

import (
	"context"
	"user/internal/api/domain/users"
)

//go:generate mockgen -source services.go -destination ./services_mock.go -package services
type (
	// Users - describe an interface for working with users
	Users interface {
		GetByName(ctx context.Context, name string) (users.User, error)
		RegisterUser(ctx context.Context, user users.User) (uint64, error)
	}
)
