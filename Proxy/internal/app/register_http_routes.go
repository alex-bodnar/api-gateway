package app

import (
	"net/http"
	"proxy/internal/config"
	"proxy/pkg/errs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/fiber/v2/utils"
)

const (
	// UserName - header used in requests.
	UserName = "Username"
)

func (a *App) registerHTTPRoutes(app *fiber.App) {
	router := app.Group("/v1/proxy")
	router.Get("/status", a.statusHTTPHandler.CheckStatus)

	for _, cfgProxy := range a.config.Proxy {
		a.registerProxyHTTPRoutes(cfgProxy, router)
	}

}

// registerProxyHTTPRoutes – registers proxy HTTP routes.
func (a *App) registerProxyHTTPRoutes(cfg config.Proxy, router fiber.Router) {
	if !cfg.Enabled {
		return
	}

	proxyGroup := router.Group(cfg.Group)
	for _, route := range cfg.Routes {
		switch route.Method {
		case http.MethodPost:
			proxyGroup.Post(route.In, a.checkAuthorization(route.CheckAuthorization), proxy.Forward(route.To))
		case http.MethodGet:
			proxyGroup.Get(route.In, a.checkAuthorization(route.CheckAuthorization), proxy.Forward(route.To))
		}
	}
}

// checkAuthorization – checks authorization.
func (a *App) checkAuthorization(required bool) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if !required {
			return c.Next()
		}

		req := c.Request()
		res := c.Response()
		originalURL := utils.CopyString(c.OriginalURL())
		defer req.SetRequestURI(originalURL)
		req.SetRequestURI(a.config.AuthorizationService.URL)

		req.Header.Del(fiber.HeaderConnection)
		if err := a.clientHTTP.Do(req, res); err != nil {
			return errs.Unauthorized{}
		}

		if res.StatusCode() != fiber.StatusOK {
			return errs.Unauthorized{}
		}

		res.Header.Del(fiber.HeaderConnection)

		return c.Next()
	}
}
