package authenticator

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"user/internal/api/domain/users"
	"user/internal/api/repository"
	"user/internal/config"
	"user/pkg/errs"
	"user/pkg/log"
)

var _ repository.Authenticator = &repo{}

// repo implements repository.Authenticator
type repo struct {
	writer *kafka.Writer
}

// NewRepository constructor.
func NewRepository(cfg config.AuthenticatorKafka, logger log.Logger) *repo {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      cfg.Brokers,
		BatchSize:    cfg.BatchSize,
		BatchTimeout: cfg.BatchTimeout,
		RequiredAcks: cfg.RequiredAcks,
		Logger:       logger,
	})

	writer.AllowAutoTopicCreation = true

	return &repo{writer: writer}
}

// SendNewUser - send new user to kafka topic.
func (r *repo) SendNewUser(ctx context.Context, user users.User) error {
	byteUser, err := json.Marshal(toDatabaseUser(user))
	if err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	msg := kafka.Message{
		Topic: registerUserTopic,
		Value: byteUser,
	}

	if err = r.writer.WriteMessages(ctx, msg); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	return nil
}
