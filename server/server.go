package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"os"
	. "schrader/server/middlewares"
	. "schrader/server/routes"
)

var (
	path   = "./conf/config.json"
	config Config
)

func init() {
	config = GetConfig(path)
}

func main() {
	e := echo.New()

	Middlewares(e)
	Routes(e)
	addr := config.Server.Host + ":" + config.Server.Port

	if config.UseSSL {
		e.Logger.Fatal(e.StartTLS(addr, config.Cert, config.Key))
	} else {
		e.Logger.Fatal(e.Start(addr))
	}
}

type Config struct {
	ServiceName    string `json:"service_name"`
	ServiceVersion string `json:"service_version"`
	Server         struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	UseSSL bool   `json:"use_ssl"`
	Cert   string `json:"cert"`
	Key    string `json:"key"`
}

func GetConfig(path string) Config {
	var config Config

	file, err := os.Open(path)
	if err != nil {
		panic("Error while opening config file (Probably not found)")
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Error while reading config file")
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic("Error while unmarshalling config file (Probably invalid JSON)")
	}

	return config
}
