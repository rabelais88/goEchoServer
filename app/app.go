package app

import (
	"goEchoServer/controller"
	"goEchoServer/router"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

func Start(run bool) http.Handler {
	godotenv.Load(".env")

	echo.NotFoundHandler = controller.NotExist
	e := echo.New()
	db := controller.ConnectDB()

	_db := controller.UseDB(db)
	e.Use(_db.SetContext)

	router.Route(e)

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
