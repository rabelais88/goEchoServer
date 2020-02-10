package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}

func RespondError(c echo.Context, err ApplicationError) error {
	return c.JSON(err.StatusCode, err)
}

func SimpleError(c echo.Context, statusCode int, code string) error {
	return c.JSON(statusCode, struct{ code string }{code: code})
}

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}

func NotExist(c echo.Context) error {
	return RespondError(c, ApplicationError{
		Message:    "page doesn't exist!",
		StatusCode: http.StatusNotFound,
		Code:       "NOT_EXIST",
	})
}
