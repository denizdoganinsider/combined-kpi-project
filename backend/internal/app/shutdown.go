package app

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func waitForShutdown(
	parentCtx context.Context,
	e *echo.Echo,
	db *sql.DB,
	logger zerolog.Logger,
) error {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	logger.Info().Msg("shutdown signal received")

	ctx, cancel := context.WithTimeout(parentCtx, 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Error().
			Err(err).
			Msg("error during server shutdown")
		return err
	}

	logger.Info().Msg("http server stopped")

	if err := db.Close(); err != nil {
		logger.Error().
			Err(err).
			Msg("error closing database connection")
		return err
	}

	logger.Info().Msg("database connection closed")
	logger.Info().Msg("graceful shutdown completed")

	return nil
}
