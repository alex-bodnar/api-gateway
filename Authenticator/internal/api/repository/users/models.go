package users

import (
	"time"

	"authenticator/internal/api/domain/users"
)

type (
	// user – user database model
	user struct {
		ID        uint64    `db:"id"`
		Name      string    `db:"name"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

// toDatabaseUser converts domain user to database user
func toDatabaseUser(u users.User) user {
	return user{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// toDomain converts database user to domain user
func (u user) toDomain() users.User {
	return users.User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
