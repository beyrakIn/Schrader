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

		response <- out

		return c.String(http.StatusOK, "Control")
	}
}
