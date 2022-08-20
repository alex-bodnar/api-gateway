package logger

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"proxy/pkg/log"
)

// Middleware is a middleware for logging HTTP requests
func Middleware(l log.Logger) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		defer func() {
			params := []interface{}{
				"duration", time.Since(start).Seconds(),
				"client_ip", c.IP(),
				"forwarded_ips", c.IPs(),
				"method", string(c.Context().Method()),
				"status_code", c.Response().StatusCode(),
				"body_size", len(c.Response().Body()),
				"path", string(c.Context().URI().Path()),
			}

			l.Infow("complete", params...)
		}()

		if err := c.Next(); err != nil {
			l.Debugf("request error: %v", err)
			return err
		}

		return nil
	}
}
