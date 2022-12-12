package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func NewClient(conn *websocket.Conn, pool *Pool, id string) *Client {
	return &Client{
		ID:   id,
		Conn: conn,
		Pool: pool,
	}
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		//c.Pool.Broadcast <- message
		fmt.Printf("%s -> %+v\n", c.ID, message.Body)
	}
}

func (c *Client) Write(messageType int, payload []byte) error {
	return c.Conn.WriteMessage(messageType, payload)
}
