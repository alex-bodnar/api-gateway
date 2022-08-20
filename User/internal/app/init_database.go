package app

import (
	"context"

	"user/pkg/database"
)

// initDatabase init database in app struct.
func (a *App) initDatabase(ctx context.Context) {
	a.db = database.InitDatabase(a.config.Storage.Postgres, a.logger, a.dbMigrationsFS)
}
