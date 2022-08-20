package users_redis

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/stdlib"
	"gotest.tools/assert"

	"authenticator/internal/api/domain/users"
	userRepo "authenticator/internal/api/repository/users_redis"
)

func TestRepository_Create(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

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
			r := userRepo.NewRepository(time.Second*100, rdb)

			if err := r.Save(tt.args.ctx, tt.args.userData); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := r.GetByName(tt.args.ctx, tt.args.userData.Name)
			if err != nil {
				t.Errorf("Repository.GetByName() error = %v", err)
			}

			assert.DeepEqual(t, got, tt.args.userData)
		})
	}
}
