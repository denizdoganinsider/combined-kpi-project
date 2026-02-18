package app

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func (a *App) setupMiddleware() {
	e := a.e

	e.IPExtractor = echo.ExtractIPFromXFFHeader()

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 4 << 10,
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			a.logger.Error().
				Err(err).
				Str("path", c.Path()).
				Bytes("stack", stack).
				Msg("panic recovered")
			return nil
		},
	}))

	e.Use(requestID)

	e.Use(a.requestLogger())

	e.Use(middleware.BodyLimit("2M"))

	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            31536000,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{a.config.FrontendOrigin},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderXRequestID,
		},
		AllowCredentials: true,
	}))

	store := middleware.NewRateLimiterMemoryStoreWithConfig(
		middleware.RateLimiterMemoryStoreConfig{
			Rate:      rate.Limit(10),
			Burst:     20,
			ExpiresIn: 3 * time.Minute,
		},
	)

	e.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Store: store,
		IdentifierExtractor: func(c echo.Context) (string, error) {
			ip := c.RealIP()
			if ip == "" {
				host, _, _ := net.SplitHostPort(c.Request().RemoteAddr)
				ip = host
			}
			return ip, nil
		},
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Path(), "/health")
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusTooManyRequests, map[string]string{
				"error": "rate limit exceeded",
			})
		},
	}))
}

func requestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Request().Header.Get(echo.HeaderXRequestID)
		if id == "" {
			id = uuid.New().String()
		}
		c.Set("request_id", id)
		c.Response().Header().Set(echo.HeaderXRequestID, id)
		return next(c)
	}
}

func (a *App) requestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			latency := time.Since(start)
			status := c.Response().Status

			log := a.logger.Info()
			if status >= 400 {
				log = a.logger.Error()
			}

			log.
				Str("method", c.Request().Method).
				Str("uri", c.Request().RequestURI).
				Int("status", status).
				Str("ip", c.RealIP()).
				Dur("latency", latency).
				Str("request_id", c.Response().Header().Get(echo.HeaderXRequestID)).
				Msg("incoming request")

			return err
		}
	}
}
