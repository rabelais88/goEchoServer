package controller

import (
	"net/http"

	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Post struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostModel struct {
	gorm.Model
	Post
}

func GetPosts(c echo.Context) error {
	cc := c.(*ServerContext)
	// cc.db.Create(&Post{Author: "TEST", Title: "TESTtitle", Content: "blahblah"})
	var post PostModel
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return SimpleError(c, http.StatusBadRequest, "NO_QUERY_PAGE")
	}
	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		return SimpleError(c, http.StatusBadRequest, "NO_QUERY_SIZE")
	}
	offset := page * size
	cc.Logger().Debug("size", size, "offset", offset, "page", page)
	cc.db.Find(&post).Offset(offset).Limit(size)
	return c.JSON(http.StatusOK, post)
}

func AddPost(c echo.Context) error {
	cc := c.(*ServerContext)
	post := new(PostModel)
	if err := cc.Bind(post); err != nil {
		return SimpleError(c, http.StatusBadRequest, "WRONG_DATA")
	}
	row := new(PostModel)
	cc.db.Create(&post).Scan(&row)
	return c.JSON(http.StatusOK, row)
}
