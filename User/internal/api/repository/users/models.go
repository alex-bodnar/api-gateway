package users

import (
	"database/sql"
	"time"

	"user/internal/api/domain/users"
	"user/internal/api/repository"
)

type (
	// usersList - list of users
	usersList []user

	// user – user database model
	user struct {
		ID          uint64         `db:"id"`
		Name        string         `db:"name"`
		Phone       string         `db:"phone"`
		DateOfBirth sql.NullTime   `db:"date_of_birth"`
		Age         sql.NullInt64  `db:"age"`
		Email       sql.NullString `db:"email"`
		CreatedAt   time.Time      `db:"created_at"`
		UpdatedAt   time.Time      `db:"updated_at"`
	}
)

// toDatabaseUser converts domain user to database user
func toDatabaseUser(u users.User) user {
	return user{
		ID:          u.ID,
		Name:        u.Name,
		Phone:       u.Phone,
		DateOfBirth: repository.ToNullTime(u.DateOfBirth),
		Age:         repository.ToNullInt64(int64(u.Age)),
		Email:       repository.ToNullString(u.Email),
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

// toDomain converts database user to domain user
func (u user) toDomain() users.User {
	return users.User{
		ID:          u.ID,
		Name:        u.Name,
		Phone:       u.Phone,
		DateOfBirth: u.DateOfBirth.Time,
		Age:         uint64(u.Age.Int64),
		Email:       u.Email.String,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

// toDomain converts list of database users to list of domain users
func (u usersList) toDomain() []users.User {
	users := make([]users.User, 0, len(u))
	for _, user := range u {
		users = append(users, user.toDomain())
	}

	return users
}
