package routes

import (
	"github.com/labstack/echo/v4"
	. "schrader/server/controllers"
	"schrader/server/pkg/websocket"
)

var (
	pool = websocket.NewPool()
)

func Routes(e *echo.Echo) {
	go pool.Start()

	e.GET("/", Run(pool))
	e.GET("/broadcast", Broadcast(pool))
	e.GET("/control", Control())
	e.GET("/client", ClientsStatus(pool))
	e.GET("client/:id", Command(pool))

}
