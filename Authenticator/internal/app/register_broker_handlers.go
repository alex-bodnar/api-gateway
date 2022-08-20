package app

import "authenticator/internal/api/delivery/broker/users"

const (
	registerUserTopic = "register-user"
)

func (a *App) brokerRoutes() map[string]MessageHandler {
	return map[string]MessageHandler{
		registerUserTopic: a.usersBrokerHandler.SaveNewUser,
	}
}

func (a *App) registerBrokerHandlers() {
	a.usersBrokerHandler = users.NewHandler(a.usersService)
}
