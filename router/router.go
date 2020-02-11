package router

import (
	"goEchoServer/controller"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, db *gorm.DB) {
	e.GET(`/`, controller.HelloWorld)
	e.GET(`*`, controller.NotExist)
	e.GET(`/posts`, controller.GetPosts(db))
}
