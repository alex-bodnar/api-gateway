package app

import (
	"context"
	"embed"

	"github.com/jmoiron/sqlx"

	"user/internal/api/delivery"
	"user/internal/api/repository"
	"user/internal/api/services"
	"user/internal/config"
	"user/pkg/http/responder"
	"user/pkg/log"
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

		responder responder.Responder

		// Repository dependencies.
		authenticatorRepo repository.Authenticator
		usersRepo         repository.Users

		// Service dependencies.
		usersService services.Users

		// Delivery dependencies.
		statusHTTPHandler delivery.StatusHTTP
		usersHTTPHandler  delivery.UsersHTTP
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

	// Run Workers
	a.runWorkers(ctx)
}
