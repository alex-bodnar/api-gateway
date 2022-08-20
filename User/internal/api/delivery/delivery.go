package delivery

import (
	"github.com/gofiber/fiber/v2"
)

type (
	// StatusHTTP – describes an interface for work with service status over HTTP.
	StatusHTTP interface {
		CheckStatus(ctx *fiber.Ctx) error
		GetName(ctx *fiber.Ctx) error
	}

	// UsersHTTP – describes an interface for work with users over HTTP.
	UsersHTTP interface {
		RegisterUser(ctx *fiber.Ctx) error
		GetUser(ctx *fiber.Ctx) error
	}
)
