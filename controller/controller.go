package controller

import (
	"net/http"
	"os"

	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// on GORM mocking
// https://github.com/jinzhu/gorm/issues/1525

func ConnectDB() *gorm.DB {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", dbHost, dbPort, dbName, dbUser, dbPass)
	fmt.Println("postgres connection :", connectionString)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Post{})
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
