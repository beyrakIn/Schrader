package controllers

import (
	"github.com/labstack/echo/v4"
	"schrader/server/pkg/websocket"
)

func Broadcast(pool *websocket.Pool) func(c echo.Context) error {
	return func(c echo.Context) error {
		// get the message from the query parameter
		message := c.QueryParam("message")
		// send the message to all clients
		pool.Broadcast <- message
		// return a status code
		return c.String(200, "Message sent to all clients -> "+message)
	}
}
