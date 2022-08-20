package users

import (
	"time"

	"user/internal/api/domain/users"
)

type (
	// userRequest - user request model
	userRequest struct {
		Name        string `json:"name"`
		Phone       string `json:"phone"`
		DateOfBirth int64  `json:"date_of_birth"`
		Email       string `json:"email"`
	}

	// registerUserResponse - response model for register user
	registerUserResponse struct {
		ID uint64 `json:"id"`
	}

	// userResponse - response models for user
	userResponse struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Phone       string `json:"phone"`
		DateOfBirth int64  `json:"date_of_birth,omitempty"`
		Age         uint64 `json:"age,omitempty"`
		Email       string `json:"email,omitempty"`
		CreatedAt   int64  `json:"created_at"`
		UpdatedAt   int64  `json:"updated_at"`
	}
)

// toDomain converts user request to domain user
func (u userRequest) toDomain() users.User {
	var dateOfBirth time.Time
	if u.DateOfBirth > 0 {
		dateOfBirth = time.Unix(u.DateOfBirth, 0)
	}

	return users.User{
		Name:        u.Name,
		Phone:       u.Phone,
		DateOfBirth: dateOfBirth,
		Email:       u.Email,
	}
}

// toResponse converts domain user to response model
func toResponseUser(u users.User) userResponse {
	var dateOfBirth int64
	if !u.DateOfBirth.IsZero() {
		dateOfBirth = u.DateOfBirth.Unix()
	}

	return userResponse{
		ID:          u.ID,
		Name:        u.Name,
		Phone:       u.Phone,
		DateOfBirth: dateOfBirth,
		Age:         u.Age,
		Email:       u.Email,
		CreatedAt:   u.CreatedAt.Unix(),
		UpdatedAt:   u.UpdatedAt.Unix(),
	}
}
