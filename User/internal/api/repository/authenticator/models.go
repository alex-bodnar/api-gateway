package authenticator

import "user/internal/api/domain/users"

const (
	// registerUserTopic – topic for notify about registration new user
	registerUserTopic = "register-user"
)

type (
	// user – user model
	user struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)

// toDatabaseUser converts domain user to database user
func toDatabaseUser(u users.User) user {
	return user{
		ID:   u.ID,
		Name: u.Name,
	}
}
