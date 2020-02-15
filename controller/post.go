package controller

import (
	"net/http"

	"strconv"

	// _ "github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"time"
)

type Post struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostModel struct {
	// overrides gorm.Model for camelcasing
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index"`
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
	cc.db.Create(&post)
	return c.JSON(http.StatusOK, post)
}

func GetPost(c echo.Context) error {
	cc := c.(*ServerContext)
	postId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return SimpleError(c, http.StatusBadRequest, "ID_NOT_NUMERIC")
	}

	post := new(PostModel)
	cc.db.First(&post, postId)
	return c.JSON(http.StatusOK, post)
}

func RemovePost(c echo.Context) error {
	cc := c.(*ServerContext)
	postId, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return SimpleError(c, http.StatusBadRequest, "ID_NOT_NUMERIC")
	}
	post := new(PostModel)
	cc.db.First(&post, postId)
	cc.db.Delete(&post)
	return c.String(http.StatusOK, "POST_DELETED")
}

func ModifyPost(c echo.Context) error {
	cc := c.(*ServerContext)
	post := new(PostModel)
	if err := cc.Bind(post); err != nil {
		return SimpleError(c, http.StatusBadRequest, "WRONG_POST_DATA")
	}
	prevPost := new(PostModel)
	cc.db.First(&prevPost)
	prevPost.Title = post.Title
	prevPost.Content = post.Content
	prevPost.Author = post.Author
	cc.db.Save(&prevPost)
	return c.String(http.StatusOK, "POST_MODIFIED")
}
