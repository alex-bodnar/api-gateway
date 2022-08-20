package app

import (
	"context"
	"embed"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"

	"authenticator/internal/api/delivery"
	"authenticator/internal/api/repository"
	"authenticator/internal/api/services"
	"authenticator/internal/config"
	"authenticator/pkg/http/responder"
	"authenticator/pkg/log"
)

type (
	// Meta defines meta for application.
	Meta struct {
		Info       Info
		ConfigPath string
	}

	// Info defines metadata of application.
	Info struct {
		AppName       string
		Tag           string
		Version       string
		Commit        string
		Date          string
		FortuneCookie string
	}

	// App defines main application struct.
	App struct {
		// meta information about application.
		meta Meta

		// tech dependencies.
		config *config.Config
		logger log.Logger

		dbMigrationsFS embed.FS
		db             *sqlx.DB
		rdb            *redis.Client

		responder responder.Responder

		// Repository dependencies.
		usersPostgresRepo repository.Users
		usersRedisRepo    repository.Users

		// Service dependencies.
		usersService services.UsersService

		// Delivery dependencies.
		statusHTTPHandler  delivery.StatusHTTP
		usersHTTPHandler   delivery.UsersHTTP
		usersBrokerHandler delivery.UsersBroker
	}

	worker func(ctx context.Context, a *App)
)

// New - app constructor without init for components.
func New(meta Meta) *App {
	return &App{
		meta: meta,
	}
}

// WithMigrationFS is a setup for database migration filesystem
func (a *App) WithMigrationFS(f embed.FS) *App {
	a.dbMigrationsFS = f
	return a
}

// Run – registers graceful shutdown.
// populate configuration and application dependencies,
// run workers.
func (a *App) Run(ctx context.Context) {
	// Initialize configuration
	a.populateConfiguration()

	// Register Dependencies
	a.initLogger()
	a.initDatabase(ctx)

	// Domain registration.
	a.registerRepositories()
	a.registerServices(ctx)

	// Register Handlers
	a.registerHTTPHandlers()
	a.registerBrokerHandlers()

	// Run Workers
	a.runWorkers(ctx)
}
