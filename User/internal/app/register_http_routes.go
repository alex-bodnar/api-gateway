package app

import "github.com/gofiber/fiber/v2"

func (a *App) registerHTTPRoutes(app *fiber.App) {
	router := app.Group("/v1/user")
	router.Get("/status", a.statusHTTPHandler.CheckStatus)

	router.Get("/microservice/name", a.statusHTTPHandler.GetName)

	user := router.Group("/user")
	user.Post("/register", a.usersHTTPHandler.RegisterUser)
	user.Get("/profile", a.usersHTTPHandler.GetUser)
}
