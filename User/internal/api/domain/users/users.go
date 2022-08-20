package users

import "time"

type (
	// User – user model
	User struct {
		ID          uint64
		Name        string
		Phone       string
		DateOfBirth time.Time
		Age         uint64
		Email       string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)

// CalculateAge - calculates user age
func (u *User) CalculateAge() {
	if u.DateOfBirth.IsZero() {
		return
	}

	u.Age = uint64(time.Now().Year() - u.DateOfBirth.Year())
}
