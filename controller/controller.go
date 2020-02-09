package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}

func Init(e *echo.Echo) {
	e.GET(`/`, helloWorld)
}
