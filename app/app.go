package app

import (
	"goEchoServer/controller"
	"goEchoServer/router"
	"net/http"
	"os"

	"github.com/friendsofgo/graphiql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(run bool) http.Handler {
	godotenv.Load(".env")

	echo.NotFoundHandler = controller.NotExist
	e := echo.New()
	db := controller.ConnectDB()

	_db := controller.UseDB(db)
	e.Use(_db.SetContext)
	e.Use(middleware.Logger())
	if run {
		e.Use(middleware.Recover())
	}

	router.Route(e)
	gqh, err := controller.GraphQLHandler(db)
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.POST("/graphql", echo.WrapHandler(gqh))
	gqih, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.GET("/graphiql", echo.WrapHandler(gqih))

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
