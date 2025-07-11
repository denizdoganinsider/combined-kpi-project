package controller

import (
	"myapp-backend/controller/request"
	"myapp-backend/controller/response"
	"myapp-backend/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MovieController struct {
	movieService *service.MovieService
}

func NewMovieController(movieService *service.MovieService) *MovieController {
	return &MovieController{
		movieService: movieService,
	}
}

func (mc *MovieController) RegisterRoutes(e *echo.Echo) {
	e.POST("/movies", mc.handleCreateMovie)
}

func (mc *MovieController) handleCreateMovie(c echo.Context) error {
	var movieReq request.Movie

	if err := c.Bind(&movieReq); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Invalid request payload",
		})
	}

	err := mc.movieService.CreateMovie(movieReq)
	if err != nil {
		return c.JSON(http.StatusConflict, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Message: "Movie created successfully",
	})
}
