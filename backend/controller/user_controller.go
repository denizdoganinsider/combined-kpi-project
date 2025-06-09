package controller

import (
	"myapp-backend/controller/request"
	"myapp-backend/controller/response"
	"myapp-backend/service"
	"myapp-backend/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userController *UserController) RegisterRoutes(e *echo.Echo) {
	e.GET("/home", userController.showHome)
	e.GET("/login", userController.showLogin)
	e.POST("/login", userController.handleLogin)
	e.POST("/signup", userController.handleSignup)
	e.RouteNotFound("/*", userController.showNotFound)
}

func (userController *UserController) showHome(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	claims, err := utils.ParseJWT(cookie.Value)
	if err != nil {
		return c.String(http.StatusSeeOther, "/login")
	}

	// Güvenli type assertion
	emailInterface := claims["email"]
	email, ok := emailInterface.(string)
	if !ok {
		return c.String(http.StatusInternalServerError, "Invalid token payload: email not string")
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"Email": email,
	})
}

func (userController *UserController) showLogin(c echo.Context) error {
	_, err := c.Cookie("token")

	if err != nil {
		return c.File("static/login.html")
	}

	return c.Redirect(http.StatusSeeOther, "/home")
}

func (userController *UserController) handleLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	return userController.authenticate(email, password, c)
}

func (userController *UserController) handleSignup(c echo.Context) error {
	name := c.FormValue("first-name")
	surname := c.FormValue("last-name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	err := userController.userService.CreateUser(request.User{
		Name:     name,
		Surname:  surname,
		Email:    email,
		Password: password,
	})

	if err != nil {
		return c.JSON(http.StatusAccepted, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return userController.authenticate(email, password, c)
}

func (userController *UserController) authenticate(email string, password string, c echo.Context) error {
	token, err := userController.userService.Authenticate(email, password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Message: err.Error(),
		})
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: false,
		Expires:  time.Now().Add(1 * time.Minute),
	})

	return c.Redirect(http.StatusSeeOther, "/home")
}

func (userController *UserController) showNotFound(c echo.Context) error {
	return c.File("static/404.html")
}
