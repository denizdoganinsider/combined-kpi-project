package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Backend is working 🚀")
	})

	e.GET("/api/v1/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Merhaba Nuxt 👋",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
