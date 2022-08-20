package users

import (
	"context"
	"errors"

	"authenticator/internal/api/domain/users"
	"authenticator/internal/api/repository"
	"authenticator/internal/api/services"
	"authenticator/pkg/errs"
	"authenticator/pkg/log"
)

var _ services.UsersService = &Service{}

// Service - defines users service struct.
type Service struct {
	UsersPostgresRepo repository.Users
	UsersRedisRepo    repository.Users

	logger log.Logger
}

// NewService - constructor.
func NewService(
	usersPostgresRepo repository.Users,
	usersRedisRepo repository.Users,

	logger log.Logger,
) *Service {
	return &Service{
		UsersPostgresRepo: usersPostgresRepo,
		UsersRedisRepo:    usersRedisRepo,

		logger: logger,
	}
}

// CheckUserByName - check user authorization by name.
func (s *Service) CheckUserByName(ctx context.Context, name string) error {
	user, err := s.UsersRedisRepo.GetByName(ctx, name)
	switch {
	case err == nil:
		return nil
	case !errors.As(err, &errs.NotFound{}):
		s.logger.Error(err)
		return errs.Unauthorized{}
	}

	user, err = s.UsersPostgresRepo.GetByName(ctx, name)
	if err != nil {
		if errors.As(err, &errs.NotFound{}) {
			s.logger.Debug(err)
			return errs.Unauthorized{}
		}

		s.logger.Error(err)
		return errs.Unauthorized{}
	}

	if err = s.UsersRedisRepo.Save(ctx, user); err != nil {
		s.logger.Error(err)
	}

	return nil
}

// SaveNewUser - save new user.
func (s *Service) SaveNewUser(ctx context.Context, userData users.User) error {
	if err := s.UsersPostgresRepo.Save(ctx, userData); err != nil {
		s.logger.Error(err)
		return errs.Internal{}
	}

	return nil
}
