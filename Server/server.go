package main

import (
	. "belial/Server/middlewares"
	. "belial/Server/routes"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

var (
	addr = ":443"
	cert = "cert/cert.crt"
	key  = "cert/cert.key"

	// colors
	green = color.Green
	red   = color.Red
	bold  = color.New(color.Bold).SprintFunc()
)

func main() {
	e := echo.New() // create new echo instance

	Middlewares(e)
	Routes(e)

	//e.Logger.Fatal(e.StartTLS(addr, cert, key))
	e.Logger.Fatal(e.Start(addr))

}
