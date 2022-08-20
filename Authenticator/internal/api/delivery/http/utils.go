package http

import (
	"github.com/gofiber/fiber/v2"

	"authenticator/pkg/errs"
)

const (
	// UserName - header used in requests.
	UserName = "Username"
)

// GetUserNameFromHeader is a helper for getting User Name from fiber context
func GetUserNameFromHeader(ctx *fiber.Ctx) (string, error) {
	userName := ctx.Get(UserName, "")
	if userName == "" {
		return "", errs.BadRequest{Cause: "Username::is_required"}
	}

	return userName, nil
}
