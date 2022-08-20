package app

import (
	"authenticator/internal/api/delivery/http/status"
	"authenticator/internal/api/delivery/http/users"
)

func (a *App) registerHTTPHandlers() {
	a.statusHTTPHandler = status.NewHandler(
		a.meta.Info.AppName,
		a.meta.Info.Tag,
		a.meta.Info.Version,
		a.meta.Info.Commit,
		a.meta.Info.Date,
		a.meta.Info.FortuneCookie,
	)
	a.usersHTTPHandler = users.NewHandler(a.usersService)
}
