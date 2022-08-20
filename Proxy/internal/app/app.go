package app

import (
	"context"

	"github.com/valyala/fasthttp"

	"proxy/internal/api/delivery"
	"proxy/internal/config"
	"proxy/pkg/http/responder"
	"proxy/pkg/log"
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

		responder  responder.Responder
		clientHTTP fasthttp.Client

		// Delivery dependencies.
		statusHTTPHandler delivery.StatusHTTP
	}

	worker func(ctx context.Context, a *App)
)

// New - app constructor without init for components.
func New(meta Meta) *App {
	return &App{
		meta: meta,
	}
}

// Run – registers graceful shutdown.
// populate configuration and application dependencies,
// run workers.
func (a *App) Run(ctx context.Context) {
	// Initialize configuration
	a.populateConfiguration()

	// Register Dependencies
	a.initLogger()

	// Register Handlers
	a.registerHTTPHandlers()

	// Run Workers
	a.runWorkers(ctx)
}
