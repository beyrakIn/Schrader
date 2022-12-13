package main

import (
	"github.com/labstack/echo/v4"
	. "schrader/server/middlewares"
	. "schrader/server/routes"
)

var (
	addr = ":443"
	cert = "cert/cert.crt"
	key  = "cert/cert.key"
)

func main() {
	e := echo.New()

	Middlewares(e)
	Routes(e)

	//e.Logger.Fatal(e.StartTLS(addr, cert, key))
	e.Logger.Fatal(e.Start(addr))
}
