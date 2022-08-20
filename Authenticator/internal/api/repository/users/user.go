package users

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"authenticator/internal/api/domain/users"
	"authenticator/internal/api/repository"
	"authenticator/pkg/errs"
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

	query := `SELECT id, name, created_at, updated_at
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

// Save - create new user.
func (r Repository) Save(ctx context.Context, userData users.User) error {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query :=
		`INSERT INTO users (
			id, name
		) VALUES (
			:id, :name
		)`

	if _, err := r.db.NamedExecContext(ctx, query, userData); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	return nil
}
