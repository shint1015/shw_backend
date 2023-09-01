package main

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"shw/api"
	"shw/utils"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/signin", func(c echo.Context) error {
		return api.SignIn(c)
	})
	e.POST("/signup", func(c echo.Context) error {
		return api.SignUp(c)
	})

	r := e.Group("api/v1/restricted")
	r.Use(echojwt.WithConfig(utils.JwtConfig))

	r.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
