package routes

import (
	. "belial/Server/controllers"
	"belial/Server/pkg/websocket"
	"github.com/labstack/echo/v4"
)

var (
	pool = websocket.NewPool()
)

func Routes(e *echo.Echo) {
	go pool.Start()

	e.GET("/", Run(pool))
	e.GET("/command", Command(pool))
	e.GET("/control", Control())
	e.GET("/client", Status(pool))
	e.GET("client/:id", Command(pool))

}

// getter for pool
func GetPool() *websocket.Pool {
	return pool
}
