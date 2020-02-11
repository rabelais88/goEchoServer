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
