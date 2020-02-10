package app

import (
	"goEchoServer/router"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

func Start(run bool) http.Handler {
	e := echo.New()
	router.Route(e)
	godotenv.Load(".env")

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
