package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Post struct {
	gorm.Model
	Author  string
	Title   string
	Content string
}

func GetPosts(c echo.Context) error {
	cc := c.(*ServerContext)
	cc.db.Create(&Post{Author: "TEST", Title: "TESTtitle", Content: "blahblah"})
	var post Post
	cc.db.First(&post, 1)
	return c.JSON(http.StatusOK, post)
}
