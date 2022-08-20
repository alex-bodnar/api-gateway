package users

import "time"

type (
	// User – user model
	User struct {
		ID        uint64
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
