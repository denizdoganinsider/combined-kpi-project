package app

import (
	"context"
	"database/sql"
	"myapp-backend/common/app"
	"myapp-backend/common/mysql"
	"myapp-backend/internal/routes"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type App struct {
	ctx    context.Context
	e      *echo.Echo
	db     *sql.DB
	logger zerolog.Logger
	config *app.ConfigurationManager
}

func New(ctx context.Context) *App {
	e := echo.New()

	config := app.NewConfigurationManager()
	db := mysql.GetConnectionPool(ctx, config.MySqlConfig)

	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()

	a := &App{
		ctx:    ctx,
		e:      e,
		db:     db,
		logger: logger,
		config: config,
	}

	a.setupMiddleware()
	routes.Register(e, db)

	return a
}

func (a *App) Run() error {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      a.e,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		a.logger.Info().Msg("server starting on :8080")

		if err := a.e.StartServer(server); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal().
				Err(err).
				Msg("server failed to start")
		}
	}()

	return waitForShutdown(a.ctx, a.e, a.db, a.logger)
}
