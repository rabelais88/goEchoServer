package controller

import (
	"net/http"

	"os"
	"strconv"

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
	limit, err := strconv.Atoi(os.Getenv("LIMIT_DEFAULT"))
	if err != nil {
		return SimpleError(c, http.StatusBadRequest, "PARSING_LIMIT_DEFAULT_ERROR")
	}
	var post Post
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return SimpleError(c, http.StatusBadRequest, "NO_QUERY_PAGE")
	}
	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		return SimpleError(c, http.StatusBadRequest, "NO_QUERY_SIZE")
	}
	offset := page * limit
	cc.Logger().debug("size", size, "offset", offset, "page", page)
	cc.db.Find(&post).Offset(offset).Limit(size)
	return c.JSON(http.StatusOK, post)
}
