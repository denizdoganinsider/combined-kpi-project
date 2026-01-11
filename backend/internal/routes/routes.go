package routes

import (
	"database/sql"
	"myapp-backend/controller"
	"myapp-backend/repository"
	"myapp-backend/service"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	controller.NewUserController(userService).RegisterRoutes(e)

	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)
	controller.NewMovieController(movieService).RegisterRoutes(e)

	transactionRepository := repository.NewTransactionRepository(db)
	balanceRepository := repository.NewBalanceRepository(db)

	transactionService := service.NewTransactionService(
		transactionRepository,
		balanceRepository,
	)

	controller.NewTransactionController(transactionService).RegisterRoutes(e)

}
