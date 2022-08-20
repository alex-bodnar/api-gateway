package users

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"gotest.tools/assert"

	"authenticator/internal/api/domain/users"
	userRepo "authenticator/internal/api/repository/users"
	"authenticator/pkg/log"
)

func TestRepository_Create(t *testing.T) {
	logger := log.New()

	DB, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=authenticator password=postgres sslmode=disable")
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
					ID:   uint64(gofakeit.Number(1, 10000000)),
					Name: gofakeit.Name(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepo.NewRepository(time.Second*100, DB)

			if err := r.Save(tt.args.ctx, tt.args.userData); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := r.GetByName(tt.args.ctx, tt.args.userData.Name)
			if err != nil {
				t.Errorf("Repository.GetByName() error = %v", err)
			}

			tt.args.userData.CreatedAt = got.CreatedAt
			tt.args.userData.UpdatedAt = got.UpdatedAt

			assert.DeepEqual(t, got, tt.args.userData)
		})
	}
}
