package app

import (
	"user/internal/api/repository/authenticator"
	"user/internal/api/repository/users"
)

func (a *App) registerRepositories() {
	a.authenticatorRepo = authenticator.NewRepository(a.config.Extra.AuthenticatorKafka, a.logger)
	a.usersRepo = users.NewRepository(a.config.Storage.Postgres.QueryTimeout, a.db)
}
