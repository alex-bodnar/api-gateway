package delivery

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

type (
	// StatusHTTP – describes an interface for work with service status over HTTP.
	StatusHTTP interface {
		CheckStatus(ctx *fiber.Ctx) error
	}

	// UsersHTTP – describes an interface for work with users over HTTP.
	UsersHTTP interface {
		CheckAuthorization(ctx *fiber.Ctx) error
	}

	// UsersBroker – describes an interface for work with users over Kafka.
	UsersBroker interface {
		SaveNewUser(ctx context.Context, msg kafka.Message)
	}
)
