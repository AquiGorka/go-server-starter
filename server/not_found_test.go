package server_test

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"os"
	"testing"
)

func TestNotFound(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:" + os.Getenv("APP_PORT"),
		Reporter: httpexpect.NewAssertReporter(t),
	})
	e.GET("/404").
		Expect().
		Status(http.StatusNotFound).Body().Equal("404")
}
