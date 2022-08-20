package app

import (
	"context"

	"user/internal/api/services/users"
)

// registerServices register services in app struct.
func (a *App) registerServices(ctx context.Context) {
	a.usersService = users.NewService(a.usersRepo, a.authenticatorRepo, a.logger)
}
