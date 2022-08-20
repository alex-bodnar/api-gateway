package users

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"gotest.tools/assert"

	"user/internal/api/domain/users"
	userRepo "user/internal/api/repository/users"
	"user/pkg/log"
)

func TestRepository_Create(t *testing.T) {
	logger := log.New()

	DB, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=users password=postgres sslmode=disable")
	if err != nil {
		logger.Fatalf("test init db failed: %s", err)
	}

	type args struct {
		ctx      context.Context
		userData users.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				userData: users.User{
					Name:        gofakeit.Name(),
					Phone:       gofakeit.Phone(),
					DateOfBirth: gofakeit.Date().Round(time.Second),
					Age:         uint64(gofakeit.Number(18, 100)),
					Email:       gofakeit.Email(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepo.NewRepository(time.Second*100, DB)

			got, err := r.Create(tt.args.ctx, tt.args.userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.args.userData.ID = got.ID
			tt.args.userData.CreatedAt = got.CreatedAt
			tt.args.userData.UpdatedAt = got.UpdatedAt

			assert.DeepEqual(t, got, tt.args.userData)
		})
	}
}

func TestRepository_GetByName(t *testing.T) {
	logger := log.New()

	DB, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=users password=postgres sslmode=disable")
	if err != nil {
		logger.Fatalf("test init db failed: %s", err)
	}

	type args struct {
		ctx      context.Context
		userData users.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				userData: users.User{
					Name:        gofakeit.Name(),
					Phone:       gofakeit.Phone(),
					DateOfBirth: gofakeit.Date().Round(time.Second),
					Age:         uint64(gofakeit.Number(18, 100)),
					Email:       gofakeit.Email(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepo.NewRepository(time.Second*100, DB)

			tt.args.userData, err = r.Create(tt.args.ctx, tt.args.userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := r.GetByName(tt.args.ctx, tt.args.userData.Name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.DeepEqual(t, got, tt.args.userData)
		})
	}
}
