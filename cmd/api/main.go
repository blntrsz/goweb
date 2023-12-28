package main

import (
	"github.com/blntrsz/goweb/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.File("/tailwind.css", "tailwind.css")
	app.GET("/", userHandler.HandleUserShow)

	app.Start(":3000")
}
