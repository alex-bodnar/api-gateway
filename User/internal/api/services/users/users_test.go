package users

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"

	"user/internal/api/domain/users"
	"user/internal/api/repository"
	"user/pkg/errs"
	"user/pkg/log"
)

func TestService_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := log.NewTest(t, zaptest.Level(zapcore.FatalLevel))

	userRepo := repository.NewMockUsers(ctrl)
	authenticatorRepo := repository.NewMockAuthenticator(ctrl)

	service := NewService(userRepo, authenticatorRepo, logger)

	type mocker struct {
		userRepo          *repository.MockUsers
		authenticatorRepo *repository.MockAuthenticator
	}

	m := mocker{
		userRepo:          userRepo,
		authenticatorRepo: authenticatorRepo,
	}

	type args struct {
		ctx  context.Context
		user users.User
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(ctx context.Context, m mocker)
		want    uint64
		wantErr bool
	}{
		{
			name: "success register user",
			args: args{
				ctx: context.Background(),
				user: users.User{
					Name:        "test user",
					Email:       "examle@mail.ua",
					Phone:       "123456789",
					DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.userRepo.EXPECT().GetByName(ctx, "test user").Return(users.User{}, errs.NotFound{})
				m.userRepo.EXPECT().Create(ctx, gomock.Any()).Return(users.User{ID: 1, Name: "test user"}, nil)
				m.authenticatorRepo.EXPECT().SendNewUser(ctx, users.User{ID: 1, Name: "test user"}).Return(nil)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Fail send new user",
			args: args{
				ctx: context.Background(),
				user: users.User{
					Name:        "test user",
					Email:       "examle@mail.ua",
					Phone:       "123456789",
					DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.userRepo.EXPECT().GetByName(ctx, "test user").Return(users.User{}, errs.NotFound{})
				m.userRepo.EXPECT().Create(ctx, gomock.Any()).Return(users.User{ID: 1, Name: "test user"}, nil)
				m.authenticatorRepo.EXPECT().SendNewUser(ctx, users.User{ID: 1, Name: "test user"}).Return(errs.Internal{})
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Fail save user",
			args: args{
				ctx: context.Background(),
				user: users.User{
					Name:        "test user",
					Email:       "examle@mail.ua",
					Phone:       "123456789",
					DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.userRepo.EXPECT().GetByName(ctx, "test user").Return(users.User{}, errs.NotFound{})
				m.userRepo.EXPECT().Create(ctx, gomock.Any()).Return(users.User{ID: 1, Name: "test user"}, errs.Internal{})
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Fail user already exists",
			args: args{
				ctx: context.Background(),
				user: users.User{
					Name:        "test user",
					Email:       "examle@mail.ua",
					Phone:       "123456789",
					DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.userRepo.EXPECT().GetByName(ctx, "test user").Return(users.User{}, nil)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Fail get user internal error",
			args: args{
				ctx: context.Background(),
				user: users.User{
					Name:        "test user",
					Email:       "examle@mail.ua",
					Phone:       "123456789",
					DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.userRepo.EXPECT().GetByName(ctx, "test user").Return(users.User{}, errs.Internal{})
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args.ctx, m)

			got, err := service.RegisterUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Service.RegisterUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := log.NewTest(t, zaptest.Level(zapcore.FatalLevel))

	userRepo := repository.NewMockUsers(ctrl)
	authenticatorRepo := repository.NewMockAuthenticator(ctrl)

	service := NewService(userRepo, authenticatorRepo, logger)

	type mocker struct {
		userRepo          *repository.MockUsers
		authenticatorRepo *repository.MockAuthenticator
	}

	m := mocker{
		userRepo:          userRepo,
		authenticatorRepo: authenticatorRepo,
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(ctx context.Context, m mocker)
		want    users.User
		wantErr bool
	}{
		{
			name: "Success get user",
			args: args{
				ctx:  context.Background(),
				name: "test user",
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.userRepo.EXPECT().GetByName(ctx, "test user").Return(users.User{ID: 1, Name: "test user"}, nil)
			},
			want:    users.User{ID: 1, Name: "test user"},
			wantErr: false,
		},
		{
			name: "Fail get user",
			args: args{
				ctx:  context.Background(),
				name: "test user",
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.userRepo.EXPECT().GetByName(ctx, "test user").Return(users.User{}, errs.Internal{})
			},
			want:    users.User{},
			wantErr: true,
		},
		{
			name: "Fail not found user",
			args: args{
				ctx:  context.Background(),
				name: "test user",
			},
			mockFn: func(ctx context.Context, m mocker) {
				m.userRepo.EXPECT().GetByName(ctx, "test user").Return(users.User{}, errs.NotFound{})
			},
			want:    users.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args.ctx, m)

			got, err := service.GetByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
