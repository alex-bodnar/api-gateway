package status

import (
	"github.com/gofiber/fiber/v2"
)

type (
	// Handler defines a handler for HTTP requests for checking status.
	Handler struct {
		resp response
	}
)

// NewHandler defines a handler constructor.
func NewHandler(appName, tag, buildVersion, buildCommit, buildDate, fortuneCookie string) *Handler {
	return &Handler{
		resp: response{
			Name:          appName,
			Version:       buildVersion,
			Tag:           tag,
			Commit:        buildCommit,
			Date:          buildDate,
			FortuneCookie: fortuneCookie,
		},
	}
}

// CheckStatus -  HTTP GET handler for status endpoint.
func (h Handler) CheckStatus(ctx *fiber.Ctx) error {
	return ctx.JSON(h.resp)
}
