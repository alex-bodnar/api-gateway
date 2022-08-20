package users

import "authenticator/internal/api/domain/users"

type (
	// user – user model
	user struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)

// toDomain - convert user model to domain model.
func (u *user) toDomain() users.User {
	return users.User{
		ID:   u.ID,
		Name: u.Name,
	}
}
