package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Control() func(c echo.Context) error {
	return func(c echo.Context) error {

		out := c.QueryParam("out")
		green(bold("User-Agent: " + c.Request().UserAgent()))
		red(out)

		response <- c.Request().RemoteAddr + " -> \n" + out

		c.Response().Header().Set("Content-Type", "text/css")
		c.Response().WriteHeader(http.StatusOK)

		return nil
	}
}
