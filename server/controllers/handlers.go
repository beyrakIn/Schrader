package controllers

import (
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	. "schrader/server/pkg/websocket"
)

var (
	receiver = make(chan string)
	sender   = make(chan string)

	// Colors
	green  = color.Green
	red    = color.Red
	blue   = color.Blue
	yellow = color.Yellow
	bold   = color.New(color.Bold).SprintFunc()
)

func Run(pool *Pool) func(c echo.Context) error {
	return func(c echo.Context) error {
		// check if request is from a websocket or not. If not, redirect request to /client.
		if c.Request().Header.Get("Upgrade") != "websocket" {
			return c.Redirect(302, "/client")
		}

		ws, err := Upgrade(c.Response(), c.Request())
		if err != nil {
			red(bold("Error: " + err.Error()))
		}

		defer ws.Close()
		green(bold("\nclient connected to server, ID: " + ws.RemoteAddr().String()))

		// new client created and added to the pool
		id := uuid.New().String()
		client := NewClient(ws, pool, id)
		pool.Register <- client
		client.Read()

		return nil
	}
}
