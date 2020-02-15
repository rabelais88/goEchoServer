package router

import (
	"goEchoServer/controller"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	e.GET(`/`, controller.HelloWorld)
	e.GET(`/post`, controller.GetPost)
	e.GET(`/posts`, controller.GetPosts)
	e.POST(`/post`, controller.AddPost)
	e.DELETE(`/post`, controller.RemovePost)
	e.PUT(`/post`, controller.ModifyPost)
}
