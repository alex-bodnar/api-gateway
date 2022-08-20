package responder

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"proxy/pkg/errs"
)

type (
	Error struct {
		Error string `json:"error"`
	}

	// Responder - helper struct wraps request/response functionality for fiber http handlers
	Responder struct{}
)

// New - constructor for Responder entity
func New() Responder {
	return Responder{}
}

// Respond - helper func for respond
func (h Responder) Respond(ctx *fiber.Ctx, code int, payload interface{}) error {
	ctx.Response().SetStatusCode(code)

	if err := ctx.JSON(payload); err != nil {
		return errs.Internal{Cause: "invalid json in output"}
	}

	return nil
}

// RespondError - helper func for response with error
func (h Responder) RespondError(ctx *fiber.Ctx, code int, err Error) error {
	return h.Respond(ctx, code, err)
}

// RespondEmpty - helper func for empty response
func (h Responder) RespondEmpty(ctx *fiber.Ctx, code int) error {
	ctx.Response().SetStatusCode(code)
	return nil
}

// HandleError - helper func for response with error
func (h Responder) HandleError(ctx *fiber.Ctx, err error) error {
	switch err := err.(type) {
	case errs.Empty:
		return h.RespondEmpty(ctx, http.StatusNoContent)
	case errs.BadGateway:
		return h.Respond(ctx, http.StatusBadGateway, Error{Error: err.Error()})
	case errs.Unauthorized:
		return h.Respond(ctx, http.StatusUnauthorized, Error{Error: err.Error()})
	case errs.Forbidden:
		return h.Respond(ctx, http.StatusForbidden, Error{Error: err.Error()})
	case errs.NotFound:
		return h.Respond(ctx, http.StatusNotFound, Error{Error: err.Error()})
	case errs.AlreadyExists, errs.Conflict:
		return h.Respond(ctx, http.StatusConflict, Error{Error: err.Error()})
	case errs.FieldsValidation, errs.BadRequest:
		return h.Respond(ctx, http.StatusBadRequest, Error{Error: err.Error()})
	default:
		return h.Respond(ctx, http.StatusInternalServerError, Error{Error: err.Error()})
	}
}
