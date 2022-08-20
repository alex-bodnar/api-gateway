package authenticator

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"

	"user/internal/api/domain/users"
	authService "user/internal/api/repository/authenticator"
	"user/internal/config"
	"user/pkg/log"
)

func Test_repo_SendNewUser(t *testing.T) {
	logger := log.New()

	cfg := config.AuthenticatorKafka{
		Brokers:      []string{"localhost:9092"},
		BatchSize:    1,
		BatchTimeout: time.Second,
		RequiredAcks: 1,
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "register-user",
		Logger:  logger,
	})

	type args struct {
		ctx  context.Context
		user users.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				user: users.User{ID: 1, Name: "test user"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := authService.NewRepository(cfg, logger)
			if err := r.SendNewUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("repo.SendNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			msg, err := reader.ReadMessage(tt.args.ctx)
			if err != nil {
				t.Errorf("reader.ReadMessage() error = %v", err)
			}

			var resp map[string]interface{}
			if err := json.Unmarshal(msg.Value, &resp); err != nil {
				t.Errorf("json.Unmarshal() error = %v", err)
			}

			id, ok := resp["id"].(float64)
			if !ok {
				t.Errorf("resp.id is not float64")
			}

			if uint64(id) != tt.args.user.ID {
				t.Errorf("resp.id = %v, want %v", id, tt.args.user.ID)
			}

			name, ok := resp["name"].(string)
			if !ok {
				t.Errorf("resp.name is not string")
			}

			if name != tt.args.user.Name {
				t.Errorf("resp.name = %v, want %v", name, tt.args.user.Name)
			}

		})
	}
}
