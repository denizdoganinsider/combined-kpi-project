package controller

import (
	"log"
	"net/http"
	"strconv"

	"myapp-backend/controller/request"
	"myapp-backend/controller/response"
	"myapp-backend/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userController *UserController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/users", userController.GetAllUsers)
	e.GET("/api/v1/users/:id", userController.GetUserById)
	e.POST("/api/v1/users", userController.AddUser)
	e.PUT("/api/v1/users/:id", userController.UpdateUsername)
	e.DELETE("/api/v1/users/:id", userController.DeleteUserById)
}

func (userController *UserController) GetAllUsers(c echo.Context) error {
	role := c.QueryParam("role")

	if len(role) == 0 {
		allUsers := userController.userService.GetAllUsers()
		return c.JSON(http.StatusOK, response.ToResponseList(allUsers))
	}

	usersWithGivenRole := userController.userService.GetUsersByRole(role)
	return c.JSON(http.StatusOK, response.ToResponseList(usersWithGivenRole))
}

func (userController *UserController) GetUserById(c echo.Context) error {
	id := c.Param("id")
	userId, _ := strconv.Atoi(id)

	userContent, err := userController.userService.GetById(int64(userId))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ToResponse(userContent))
}

func (userController *UserController) AddUser(c echo.Context) error {
	if c.Request().ContentLength == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Request body is required",
		})
	}

	if c.Request().Header.Get(echo.HeaderContentType) != echo.MIMEApplicationJSON {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Content-Type must be application/json",
		})
	}

	var addUserRequest request.AddUserRequest

	bindError := c.Bind(&addUserRequest)

	log.Println("Bind Error: ", bindError)

	if bindError != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: bindError.Error(),
		})
	}

	validationError := userController.userService.AddUser(addUserRequest.ToModel())

	if validationError != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			Message: validationError.Error(),
		})
	}

	return c.NoContent(http.StatusCreated)
}

func (userController *UserController) UpdateUsername(c echo.Context) error {
	id := c.Param("id")
	userId, _ := strconv.Atoi(id)

	username := c.QueryParam("username")

	if len(username) == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Username is required",
		})
	}

	if len(username) < 4 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "Username should have at least 4 characters",
		})
	}

	userController.userService.UpdateUsername(username, int64(userId))

	return c.NoContent(http.StatusOK)
}

func (userController *UserController) DeleteUserById(c echo.Context) error {
	id := c.Param("id")
	userId, _ := strconv.Atoi(id)

	err := userController.userService.DeleteById(int64(userId))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}
