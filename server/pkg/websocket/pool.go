package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan string
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan string),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Printf("client %s joined \nSize of Connection Pool: %d\n", client.ID, len(pool.Clients))
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println()
			fmt.Printf("client %s joined \nSize of Connection Pool: %d\n", client.ID, len(pool.Clients))
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client, _ := range pool.Clients {
				client.Write(websocket.TextMessage, []byte(message))
			}
		}
	}
}
