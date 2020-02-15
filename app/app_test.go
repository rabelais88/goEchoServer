package app

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v4"
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
	// t.Log(reqWrite)
	reqWrite.Object().ContainsKey("id").ValueEqual("author", post.Author).ValueEqual("title", post.Title)

	reqGet := e.GET("/post").WithQuery("id", 1).Expect().Status(http.StatusOK).JSON()
	// t.Log(reqGet)
	reqGet.Object().ContainsKey("id").ValueEqual("author", "KIM").ValueEqual("title", "mystery of orient express")

	gofakeit.Seed(0)
	sampleSize := 20

	for i := 0; i < sampleSize; i += 1 {
		_post := postForm{
			Author:  gofakeit.Name(),
			Title:   gofakeit.Sentence(15),
			Content: gofakeit.Paragraph(2, 5, 10, " "),
		}
		e.POST("/post").WithForm(_post).Expect().Status(http.StatusOK)
	}

	type pagingQuery struct {
		Size string `url:"size"`
		Page string `url:"page"`
	}
	e.GET("/posts").Expect().Status(http.StatusBadRequest)
	query := pagingQuery{
		Size: "10",
		Page: "0",
	}
	reqGetList := e.GET("/posts").WithQueryObject(query).Expect().Status(http.StatusOK).JSON()
	reqGetList.Object().ValueEqual("count", 21).ValueEqual("page", 0).Value("items").Array().Length().Equal(10)

	e.DELETE("/post").WithQuery("id", 1).Expect().Status(http.StatusOK)
	reqGetDeletedList := e.GET("/posts").WithQueryObject(query).Expect().Status(http.StatusOK).JSON()
	reqGetDeletedList.Object().ValueEqual("count", 20)
}
