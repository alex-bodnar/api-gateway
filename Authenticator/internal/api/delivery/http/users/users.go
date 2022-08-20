package users

import (
	"github.com/gofiber/fiber/v2"

	"authenticator/internal/api/delivery/http"
	"authenticator/internal/api/services"
	"authenticator/pkg/http/responder"
)

// Handler - define http handler struct for handling users requests.
type Handler struct {
	responder.Responder
	usersService services.UsersService
}

// NewHandler - constructor.
func NewHandler(usersService services.UsersService) *Handler {
	return &Handler{
		usersService: usersService,
	}
}

// CheckAuthorization - check user authorization by name.
func (h *Handler) CheckAuthorization(ctx *fiber.Ctx) error {
	userName, err := http.GetUserNameFromHeader(ctx)
	if err != nil {
		return err
	}

	if err = h.usersService.CheckUserByName(ctx.Context(), userName); err != nil {
		return err
	}

	return h.RespondEmpty(ctx, fiber.StatusOK)
}
