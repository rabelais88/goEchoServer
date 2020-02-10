package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Post struct {
	gorm.Model
	Author string
	Title string
	Content string
}

func GetPosts(db *gorm.DB) echo.HandlerFunc {
	return func (c echo.Context) error {
		db.AutoMigrate(&Post{})
		db.Create(&Post{ Author: "TEST", Title: "TESTtitle", Content: "blahblah" })
		var post Post
		db.First(&post, 1)
		return c.JSON(http.StatusOK, post)
	}
}