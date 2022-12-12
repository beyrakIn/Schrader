package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Middlewares(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	//e.Use(middleware.Logger())
	//e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{}))

}
