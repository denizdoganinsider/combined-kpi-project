package main

import (
	"context"
	"fmt"
	"log"
	"myapp-backend/common/app"
	"myapp-backend/common/mysql"
	"myapp-backend/controller"
	"myapp-backend/repository"
	"myapp-backend/response"
	"myapp-backend/service"
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
				return c.JSON(403, response.ErrorResponse{
					ErrorDescription: "Forbidden",
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

	configurationManager := app.NewConfigurationManager()

	db := mysql.GetConnectionPool(ctx, configurationManager.MySqlConfig)

	if db == nil {
		log.Fatalf("Error: Database connection is nil")
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(e)

	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)
	movieController := controller.NewMovieController(movieService)
	movieController.RegisterRoutes(e)

	/* Graceful shutdown handling */
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := e.Start(":8080"); err != nil {
			log.Fatalf("Error starting Echo server: %v", err)
		}
	}()

	<-sigs

	// Shutdown server
	fmt.Println("Shutting down server...")
	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	if err := db.Close(); err != nil {
		log.Fatalf("Error closing database: %v", err)
	}
	fmt.Println("Database connection closed.")
}
