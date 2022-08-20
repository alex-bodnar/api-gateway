package users

import (
	"authenticator/internal/api/services"
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

// Handler - define broker handler struct for handling users requests.
type Handler struct {
	usersService services.UsersService
}

// NewHandler - constructor.
func NewHandler(usersService services.UsersService) *Handler {
	return &Handler{
		usersService: usersService,
	}
}

// SaveNewUser - save new user to database.
func (h Handler) SaveNewUser(ctx context.Context, msg kafka.Message) {
	var req user
	if err := json.Unmarshal(msg.Value, &req); err != nil {
		return
	}

	h.usersService.SaveNewUser(ctx, req.toDomain())
}
