package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// on GORM mocking
// https://github.com/jinzhu/gorm/issues/1525

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	return db
}

type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}

func RespondError(c echo.Context, err ApplicationError) error {
	return c.JSON(err.StatusCode, err)
}

func SimpleError(c echo.Context, statusCode int, code string) error {
	errorBody := struct{ code string }{code: code}
	return c.JSON(statusCode, errorBody)
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
