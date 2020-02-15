package app

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func setup() {
	os.Setenv("TEST", "true")
}

// test example: https://github.com/gavv/httpexpect/blob/master/_examples/echo_test.go
func TestStart(t *testing.T) {
	setup()
	handler := Start(false)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{httpexpect.NewDebugPrinter(t, true)},
	})

	e.GET("/").Expect().Status(http.StatusOK)
}

func TestPosts(t *testing.T) {
	setup()
	handler := Start(false)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{httpexpect.NewDebugPrinter(t, true)},
	})

	type pagingQuery struct {
		Size string `url:"size"`
		Page string `url:"page"`
	}
	e.GET("/posts").Expect().Status(http.StatusBadRequest)
	query := pagingQuery{
		Size: "10",
		Page: "10",
	}
	req := e.GET("/posts").WithQueryObject(query).Expect().Status(http.StatusOK).JSON()
	t.Log(req)

	type postForm struct {
		Author  string `form:"author"`
		Title   string `form:"title"`
		Content string `form:"content"`
	}

	post := postForm{
		Author:  "KIM",
		Title:   "mystery of orient express",
		Content: "this is a good book",
	}

	reqWrite := e.POST("/post").WithForm(post).Expect().Status(http.StatusOK).JSON()
	t.Log(reqWrite)

	reqGet := e.GET("/post").WithQuery("id", 1).Expect().Status(http.StatusOK).JSON()
	t.Log(reqGet)
}
