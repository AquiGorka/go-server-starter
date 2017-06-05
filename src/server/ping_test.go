package server_test

import(
  "testing"
  "github.com/gavv/httpexpect"
  "net/http"
  "os"
)

func TestPing(t *testing.T) {
  e := httpexpect.WithConfig(httpexpect.Config{
      BaseURL: "http://localhost:" + os.Getenv("APP_PORT"),
      Reporter: httpexpect.NewAssertReporter(t),
  })
  e.GET("/ping").
    Expect().
    Status(http.StatusOK).Body().Equal("pong")
}
