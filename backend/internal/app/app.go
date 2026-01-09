package app

import (
	"context"
	"database/sql"
	"myapp-backend/common/app"
	"myapp-backend/common/mysql"
	"myapp-backend/internal/routes"

	"github.com/labstack/echo/v4"
)

type App struct {
	ctx context.Context
	e   *echo.Echo
	db  *sql.DB
}

func New(ctx context.Context) *App {
	e := echo.New()

	config := app.NewConfigurationManager()
	db := mysql.GetConnectionPool(ctx, config.MySqlConfig)

	a := &App{
		ctx: ctx,
		e:   e,
		db:  db,
	}

	a.setupMiddleware()
	routes.Register(e, db)

	return a
}

func (a *App) Run() error {
	go a.e.Start(":8080")
	return waitForShutdown(a.ctx, a.e, a.db)
}
