package main

import (
	"MVC/configs"
	"MVC/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDB()
	e := echo.New()
	routes.New(e)

	e.Logger.Fatal(e.Start(":8000"))
}
