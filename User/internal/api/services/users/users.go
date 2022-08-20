package users

import (
	"context"
	"errors"

	"user/internal/api/domain/users"
	"user/internal/api/repository"
	"user/internal/api/services"
	"user/pkg/errs"
	"user/pkg/log"
)

var _ services.Users = &Service{}

// Service - defines user service structure.
type Service struct {
	userRepo          repository.Users
	authenticatorRepo repository.Authenticator

	logger log.Logger
}

// NewService - constructor for user service.
func NewService(
	userRepo repository.Users,
	authenticatorRepo repository.Authenticator,
	logger log.Logger,
) *Service {
	return &Service{
		userRepo:          userRepo,
		authenticatorRepo: authenticatorRepo,
		logger:            logger,
	}
}

// RegisterUser - register new user.
func (s Service) RegisterUser(ctx context.Context, user users.User) (uint64, error) {
	_, err := s.userRepo.GetByName(ctx, user.Name)
	switch {
	case err == nil:
		s.logger.Debugf("user with name %s already exists", user.Name)
		return 0, errs.AlreadyExists{What: user.Name}
	case !errors.As(err, &errs.NotFound{}):
		s.logger.Error(err)
		return 0, errs.Internal{}
	}

	user.CalculateAge()

	newUser, err := s.userRepo.Create(ctx, user)
	if err != nil {
		s.logger.Error(err)
		return 0, errs.Internal{}
	}

	if err = s.authenticatorRepo.SendNewUser(ctx, newUser); err != nil {
		s.logger.Error(err)
		return 0, errs.Internal{}
	}

	return newUser.ID, nil
}

// GetByName - get user by name.
func (s Service) GetByName(ctx context.Context, name string) (users.User, error) {
	user, err := s.userRepo.GetByName(ctx, name)
	if err != nil {
		if errors.As(err, &errs.NotFound{}) {
			s.logger.Debug(err)
			return users.User{}, errs.NotFound{What: "user"}
		}

		s.logger.Error(err)
		return users.User{}, errs.Internal{}
	}

	return user, nil
}
