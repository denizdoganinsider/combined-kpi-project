package main

import (
	"log"
	"myapp-backend/controller"
	"myapp-backend/db"
	"myapp-backend/repository"
	"myapp-backend/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	db := db.GetConnection()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status}\n",
	}))

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(e)

	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)
	movieController := controller.NewMovieController(movieService)
	movieController.RegisterRoutes(e)

	log.Fatalln(e.Start(":3000"))
}
