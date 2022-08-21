package users

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"

	"authenticator/internal/api/domain/users"
	"authenticator/internal/api/repository"
	"authenticator/pkg/errs"
	"authenticator/pkg/log"
)

func TestService_CheckUserByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := log.NewTest(t, zaptest.Level(zapcore.FatalLevel))

	UsersPostgresRepo := repository.NewMockUsers(ctrl)
	UsersRedisRepo := repository.NewMockUsers(ctrl)

	service := NewService(UsersPostgresRepo, UsersRedisRepo, logger)

	type mocker struct {
		UsersRedisRepo    *repository.MockUsers
		UsersPostgresRepo *repository.MockUsers
	}

	m := mocker{
		UsersRedisRepo:    UsersRedisRepo,
		UsersPostgresRepo: UsersPostgresRepo,
	}

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(ctx context.Context, m mocker)
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				name: "test name",
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.UsersRedisRepo.EXPECT().GetByName(ctx, "test name").Return(users.User{}, errs.NotFound{})
				m.UsersPostgresRepo.EXPECT().GetByName(ctx, "test name").Return(users.User{Name: "test name"}, nil)
				m.UsersRedisRepo.EXPECT().Save(ctx, users.User{Name: "test name"}).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "fail save to redis",
			args: args{
				ctx:  context.Background(),
				name: "test name",
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.UsersRedisRepo.EXPECT().GetByName(ctx, "test name").Return(users.User{}, errs.NotFound{})
				m.UsersPostgresRepo.EXPECT().GetByName(ctx, "test name").Return(users.User{Name: "test name"}, nil)
				m.UsersRedisRepo.EXPECT().Save(ctx, users.User{Name: "test name"}).Return(errs.Internal{})
			},
			wantErr: false,
		},
		{
			name: "fail get from postgres",
			args: args{
				ctx:  context.Background(),
				name: "test name",
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.UsersRedisRepo.EXPECT().GetByName(ctx, "test name").Return(users.User{}, errs.NotFound{})
				m.UsersPostgresRepo.EXPECT().GetByName(ctx, "test name").Return(users.User{}, errs.Internal{})
			},
			wantErr: true,
		},
		{
			name: "fail not found in postgres",
			args: args{
				ctx:  context.Background(),
				name: "test name",
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.UsersRedisRepo.EXPECT().GetByName(ctx, "test name").Return(users.User{}, errs.NotFound{})
				m.UsersPostgresRepo.EXPECT().GetByName(ctx, "test name").Return(users.User{}, errs.NotFound{})
			},
			wantErr: true,
		},
		{
			name: "fail get from redis",
			args: args{
				ctx:  context.Background(),
				name: "test name",
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.UsersRedisRepo.EXPECT().GetByName(ctx, "test name").Return(users.User{}, errs.Internal{})
			},
			wantErr: true,
		},
		{
			name: "success get from redis",
			args: args{
				ctx:  context.Background(),
				name: "test name",
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.UsersRedisRepo.EXPECT().GetByName(ctx, "test name").Return(users.User{}, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args.ctx, m)

			if err := service.CheckUserByName(tt.args.ctx, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Service.CheckUserByName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_SaveNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := log.NewTest(t, zaptest.Level(zapcore.FatalLevel))

	UsersPostgresRepo := repository.NewMockUsers(ctrl)
	UsersRedisRepo := repository.NewMockUsers(ctrl)

	service := NewService(UsersPostgresRepo, UsersRedisRepo, logger)

	type mocker struct {
		UsersRedisRepo    *repository.MockUsers
		UsersPostgresRepo *repository.MockUsers
	}

	m := mocker{
		UsersRedisRepo:    UsersRedisRepo,
		UsersPostgresRepo: UsersPostgresRepo,
	}

	type args struct {
		ctx      context.Context
		userData users.User
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(ctx context.Context, m mocker)
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				userData: users.User{ID: 1, Name: "test name"},
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.UsersPostgresRepo.EXPECT().Save(ctx, users.User{ID: 1, Name: "test name"}).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "fail save to postgres",
			args: args{
				ctx:      context.Background(),
				userData: users.User{ID: 1, Name: "test name"},
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.UsersPostgresRepo.EXPECT().Save(ctx, users.User{ID: 1, Name: "test name"}).Return(errs.Internal{})
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args.ctx, m)

			if err := service.SaveNewUser(tt.args.ctx, tt.args.userData); (err != nil) != tt.wantErr {
				t.Errorf("Service.SaveNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
