package controllers

import (
	"belial/Server/pkg/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Status(pool *websocket.Pool) func(c echo.Context) error {
	return func(c echo.Context) error {
		var clients []string
		for client, _ := range pool.Clients {
			clients = append(clients, client.ID)
			c.HTML(http.StatusOK, "<a href='/client/"+client.ID+"'>"+client.ID+"</a></br>")
		}

		return nil
	}
}
