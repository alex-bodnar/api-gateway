package app

import (
	"authenticator/internal/api/repository/users"
	"authenticator/internal/api/repository/users_redis"
)

func (a *App) registerRepositories() {
	a.usersPostgresRepo = users.NewRepository(a.config.Storage.Postgres.QueryTimeout, a.db)
	a.usersRedisRepo = users_redis.NewRepository(a.config.Extra.RedisCache.TimeLive, a.rdb)
}
