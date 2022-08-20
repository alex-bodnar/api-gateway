package app

import (
	"context"

	"authenticator/internal/api/services/users"
)

// registerServices register services in app struct.
func (a *App) registerServices(ctx context.Context) {
	a.usersService = users.NewService(a.usersPostgresRepo, a.usersRedisRepo, a.logger)
}
