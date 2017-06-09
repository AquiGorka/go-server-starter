package server_test

import (
  "testing"
  "os"
  "gopkg.in/kataras/iris.v6"
  "github.com/AquiGorka/go-server-starter/src/server"
  "golang.org/x/net/websocket"
  "net/http"
  "context"
)

func TestMain(m *testing.M) {
  // run the server
  app := iris.New()
  app = server.HttpServer(app)
  app = server.WebsocketServer(app)
  go app.Listen(":" + os.Getenv("APP_PORT"))
  // tests
  code := m.Run()
  // stop the server
  ctx, _ := context.WithCancel(context.Background())
  app.Shutdown(ctx)
  // fin
  os.Exit(code)
}

func TestSocketServer(t *testing.T) {
  _, err := websocket.Dial("ws://localhost:" + os.Getenv("APP_PORT")  + "/ws", "", "http://localhost/")
  if (err != nil) {
    t.Error("Error connecting to socket server: ", err)
  }
}

func TestHttpServer(t *testing.T) {
  _, err := http.Get("http://localhost:" + os.Getenv("APP_PORT") + "/ping")
  if (err != nil) {
    t.Error("Error connecting to http server: ", err)
  }
}
