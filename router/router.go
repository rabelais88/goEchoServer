package router

import (
	"goEchoServer/controller"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	e.GET(`/`, controller.HelloWorld)
	e.GET(`*`, controller.NotExist)
}
