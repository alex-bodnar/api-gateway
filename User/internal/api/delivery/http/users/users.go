package users

import (
	"github.com/gofiber/fiber/v2"

	"user/internal/api/delivery/http"
	"user/internal/api/services"
	"user/pkg/errs"
	"user/pkg/http/responder"
)

// Handler - define http handler struct for handling users requests.
type Handler struct {
	responder.Responder
	usersService services.Users
}

// NewHandler - constructor.
func NewHandler(usersService services.Users) *Handler {
	return &Handler{
		usersService: usersService,
	}
}

// RegisterUser - register new user.
func (h *Handler) RegisterUser(ctx *fiber.Ctx) error {
	var err error

	var req userRequest
	if err := ctx.BodyParser(&req); err != nil {
		return errs.BadRequest{Cause: "invalid body"}
	}

	if errsList := req.Validate(); len(errsList) != 0 {
		return errs.FieldsValidation{Errors: errsList}
	}

	var resp registerUserResponse
	resp.ID, err = h.usersService.RegisterUser(ctx.Context(), req.toDomain())
	if err != nil {
		return err
	}

	return h.Respond(ctx, fiber.StatusCreated, resp)
}

// GetUser - get user by name.
func (h *Handler) GetUser(ctx *fiber.Ctx) error {
	userName, err := http.GetUserNameFromHeader(ctx)
	if err != nil {
		return err
	}

	resp, err := h.usersService.GetByName(ctx.Context(), userName)
	if err != nil {
		return err
	}

	return h.Respond(ctx, fiber.StatusOK, toResponseUser(resp))
}
