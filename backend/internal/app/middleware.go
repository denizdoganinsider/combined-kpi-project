package app

import (
	"myapp-backend/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func (a *App) setupMiddleware() {
	e := a.e

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(100))))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.Use(originGuard)

	e.Use(requestID)

	e.Use(middleware.Logger())
}

func originGuard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		origin := c.Request().Header.Get("Origin")
		if origin != "" && origin != "http://127.0.0.1:8080" {
			return c.JSON(403, response.ErrorResponse{ErrorDescription: "Forbidden"})
		}
		return next(c)
	}
}

func requestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Request().Header.Get(echo.HeaderXRequestID)
		if id == "" {
			id = uuid.New().String()
		}
		c.Response().Header().Set(echo.HeaderXRequestID, id)
		return next(c)
	}
}
