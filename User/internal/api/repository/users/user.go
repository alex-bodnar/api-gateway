package users

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"user/internal/api/domain/users"
	"user/internal/api/repository"
	"user/pkg/errs"
)

var _ repository.Users = &Repository{}

// Repository implements repository.Users.
type Repository struct {
	queryTimeout time.Duration
	db           *sqlx.DB
}

// NewRepository constructor.
func NewRepository(qt time.Duration, db *sqlx.DB) *Repository {
	return &Repository{
		queryTimeout: qt,
		db:           db,
	}
}

// GetByName - get user by name.
func (r Repository) GetByName(ctx context.Context, name string) (users.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query := `SELECT id, name, phone, date_of_birth, age, email, created_at, updated_at
			  FROM users WHERE name = $1`

	var result user
	err := r.db.GetContext(ctx, &result, query, name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return users.User{}, errs.Internal{Cause: err.Error()}
		}

		return users.User{}, errs.NotFound{What: "user"}
	}

	return result.toDomain(), nil
}

// Create - create new user.
func (r Repository) Create(ctx context.Context, userData users.User) (users.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query :=
		`INSERT INTO users (
			name, phone, date_of_birth, age, email
		) VALUES (
			:name, :phone, :date_of_birth, :age, :email
		) RETURNING
			id, name, phone, date_of_birth, age, email, created_at, updated_at`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return users.User{}, errs.Internal{Cause: err.Error()}
	}

	defer func() {
		_ = stmt.Close()
	}()

	var result user
	if err = stmt.GetContext(ctx, &result, toDatabaseUser(userData)); err != nil {
		return users.User{}, errs.Internal{Cause: err.Error()}
	}

	return result.toDomain(), nil
}
