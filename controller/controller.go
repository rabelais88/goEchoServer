package controller

import (
	"net/http"
	"os"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
)

// on GORM mocking
// https://github.com/jinzhu/gorm/issues/1525

type ServerContext struct {
	echo.Context
	db *gorm.DB
}

type withDB struct {
	db *gorm.DB
}

func UseDB(db *gorm.DB) *withDB {
	return &withDB{db: db} // 이미 포인터 상태로 받았기 때문에 그대로 전달해도 된다.
}

func (db *withDB) SetContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &ServerContext{c, db.db}
		return next(cc)
	}
}

func ConnectDB() *gorm.DB {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbSSL := os.Getenv("POSTGRES_SSL")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", dbHost, dbPort, dbName, dbUser, dbPass)
	if dbSSL == "disable" {
		connectionString += " sslmode=disable"
	}
	driver := "postgres"
	if os.Getenv("TEST") == "true" {
		driver = "sqlite3"
		connectionString = ":memory:"
	}
	db, err := gorm.Open(driver, connectionString)
	if err != nil {
		fmt.Println("GORM error:", err)
		panic("failed to connect database")
	}
	// defer db.Close()
	initDB(db)
	return db
}

func initDB(db *gorm.DB) {
	db.AutoMigrate(&Post{})
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
