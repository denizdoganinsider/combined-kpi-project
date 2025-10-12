package controller

import (
	"net/http"

	"myapp-backend/controller/request"
	"myapp-backend/controller/response"
	"myapp-backend/service"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	userService service.IUserService
}

func NewAuthController(userService service.IUserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (authController *AuthController) RegisterRoutes(e *echo.Echo) {
	e.POST("/api/v1/auth/register", authController.Register)
	e.POST("/api/v1/auth/login", authController.Login)
}

func (authController *AuthController) Register(c echo.Context) error {
	var req request.AddUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	user := req.ToModel()
	if err := authController.userService.AddUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse{Message: "User registered successfully"})
}

func (authController *AuthController) Login(c echo.Context) error {
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	token, err := authController.userService.Authenticate(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, response.TokenResponse{Token: token})
}
