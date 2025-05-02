package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(100))))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:8080"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			origin := c.Request().Header.Get("Origin")

			if origin != "" && origin != "http://127.0.0.1:8080" && origin != "http://localhost:3000" {
				return c.JSON(403, map[string]string{
					"ErrorDescription": "Forbidden",
				})
			}
			return next(c)
		}
	})

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${method} ${uri} ${status} ${latency_human}\n",
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get(echo.HeaderXRequestID)
			if requestID == "" {
				requestID = uuid.New().String()
				c.Request().Header.Set(echo.HeaderXRequestID, requestID)
			}
			c.Response().Header().Set(echo.HeaderXRequestID, requestID)
			return next(c)
		}
	})

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[${time_rfc3339}] ${id} ${method} ${uri} ${status} ${latency_human}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Backend is working")
	})

	e.GET("/api/v1/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Merhaba Nuxt 👋",
		})
	})

	// Graceful shutdown handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	// Start echo server and it runs background
	go func() {
		if err := e.Start(":8080"); err != nil {
			log.Fatalf("Error starting Echo server: %v", err)
		}
	}()

	// Wait for termination signal
	<-sigs

	// Shutdown server
	fmt.Println("Shutting down server...")
	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
}
