package app

import (
	"net/http"
	"os"

	"github.com/rabelais88/goEchoServer/controller"
	"github.com/rabelais88/goEchoServer/router"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

func Start(run bool) http.Handler {
	godotenv.Load(".env")

	e := echo.New()
	db := controller.ConnectDB()
	router.Route(e, db)

	if run {
		port := ":" + os.Getenv("PORT")
		e.Logger.Debug(`port`, port)

		err := e.Start(port)
		if err != nil {
			e.Logger.Fatal(err)
		}

	}
	return e
}
