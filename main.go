package main

import (
	"goEchoServer/controller"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port := ":" + os.Getenv("PORT")
	if port != ":" {
		port = ":4500"
	}

	e := echo.New()
	controller.Init(e)

	// e.GET(`/`, func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, world!")
	// })

	err := e.Start(port)
	if err != nil {
		e.Logger.Fatal(err)
	}
}
