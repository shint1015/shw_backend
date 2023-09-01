package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"shw/api"
)

func Init() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(api.MethodOverride)
	e.GET("/signup", api.SignUp)
	e.GET("/signin", api.SignIn)
	e.GET("/signout", api.SignOut)
	e.GET("/restricted", api.Restricted)
	e.Logger.Fatal(e.Start(":8080"))
}
