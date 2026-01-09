package app

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
)

func waitForShutdown(ctx context.Context, e *echo.Echo, db *sql.DB) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	e.Shutdown(ctx)
	db.Close()
	return nil
}
