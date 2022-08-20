package app

import (
	"context"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	mwLogger "proxy/pkg/http/middleware/logger"
)

func serveHTTP(ctx context.Context, app *App) {
	router := fiber.New(fiber.Config{
		Prefork:      false,
		ReadTimeout:  app.config.Delivery.HTTPServer.ReadTimeout,
		WriteTimeout: app.config.Delivery.HTTPServer.WriteTimeout,
		Network:      fiber.NetworkTCP4,
		BodyLimit:    app.config.Delivery.HTTPServer.BodySizeLimitBytes,
		AppName:      app.meta.Info.AppName,
		ErrorHandler: app.responder.HandleError,
	})

	router.Use(requestid.New())
	router.Use(recover.New())

	if app.config.Delivery.HTTPServer.LogRequests {
		router.Use(mwLogger.Middleware(app.logger))
	}

	app.registerHTTPRoutes(router)

	// graceful shutdown listener.
	go func() {
		<-ctx.Done()

		if err := router.Shutdown(); err != nil {
			app.logger.Info("🔵 http: server shutdown: %v", err)
		}
	}()

	// starting server
	ip := app.config.Delivery.HTTPServer.ListenAddress
	if err := router.Listen(ip); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			app.logger.Fatalf("🔴 failed to start server: %v", err)
		}
	}
}
