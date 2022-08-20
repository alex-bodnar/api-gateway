package app

import (
	"context"

	"github.com/segmentio/kafka-go"

	"authenticator/internal/config"
)

// MessageHandler function for handling nats message
type MessageHandler func(ctx context.Context, msg kafka.Message)

// serveBroker listen for registered subjects.
func serveBroker(ctx context.Context, app *App) {
	routes := app.brokerRoutes()

	for s, handler := range routes {
		app.logger.Infof("subscribe to: %q", s)

		subscribe(ctx, app.meta.Info.AppName, app.config.Delivery.KafkaBroker, s, handler, app.logger)
	}
}

// subscribe to subject
func subscribe(
	ctx context.Context,
	appName string,
	cfg config.KafkaBroker,
	topic string,
	handler MessageHandler,
	logger kafka.Logger,
) {
	go func() {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers: cfg.Brokers,
			GroupID: appName,
			Topic:   topic,
			Logger:  logger,
		})

		defer reader.Close()

		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				return
			}

			go handler(ctx, m)
		}
	}()
}
