package controllers

import (
	. "belial/Server/pkg/websocket"
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	p        interface{}
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
		p = pool
		ws, err := Upgrade(c.Response(), c.Request())
		if err != nil {
			red(bold("Error: " + err.Error()))
		}

		defer ws.Close()
		green(bold("Client connected to server, ID: " + ws.RemoteAddr().String()))

		// Create a new client
		id := ws.RemoteAddr().String()
		client := NewClient(ws, pool, id)
		client.ID = id
		pool.Register <- client
		fmt.Printf("Client %s registered to pool", client.ID)
		client.Read()

		//go receiveMessage(ws)

		//for {
		//	select {
		//	case this := <-sender:
		//		fmt.Println(this)
		//		//sendMessage(ws, this.code)
		//	case this := <-receiver:
		//		yellow("recv: %s", this)
		//	}
		//}

		return nil

	}
}

// function to sendMessage to server
func sendMessage(conn *websocket.Conn, message string) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		red(bold("Error: " + err.Error()))
		return
	}
}

// function to receiveMessage from server
func receiveMessage(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			red(bold("Error: " + err.Error()))
			return
		}
		receiver <- string(message)
	}
}
