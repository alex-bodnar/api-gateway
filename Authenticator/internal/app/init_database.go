package app

import (
	"context"

	"authenticator/pkg/database"
	"authenticator/pkg/redis"
)

// initDatabase init database in app struct.
func (a *App) initDatabase(ctx context.Context) {
	a.db = database.InitDatabase(a.config.Storage.Postgres, a.logger, a.dbMigrationsFS)
	a.rdb = redis.InitRedis(ctx, a.config.Storage.Redis, a.logger)
}
