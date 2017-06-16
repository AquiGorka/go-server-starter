package server_test

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"os"
	"testing"
)

func TestPing(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:" + os.Getenv("APP_PORT"),
		Reporter: httpexpect.NewAssertReporter(t),
	})
	e.GET("/ping").
		Expect().
		Status(http.StatusOK).Body().Equal("pong")
}
