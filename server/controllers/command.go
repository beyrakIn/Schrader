package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"schrader/server/pkg/websocket"
)

var (
	response = make(chan string)
)

func Command(pool *websocket.Pool) func(c echo.Context) error {
	return func(c echo.Context) error {
		var client *websocket.Client
		id := c.Param("id")
		isFound := false

		for v, _ := range pool.Clients {
			if v.ID == id {
				isFound = true
				client = v
			}
		}

		if !isFound {
			return c.String(http.StatusNotFound, "client "+id+" not found. \nBe sure to use the correct ID")
		}

		if c.QueryParam("cmd") != "" {
			err := client.Conn.WriteMessage(1, []byte(c.QueryParam("cmd")))
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			for {
				select {
				case this := <-response:
					return c.String(http.StatusOK, this)
				}
			}
		}

		return c.JSON(http.StatusNotFound, "Command not found")
	}
}
