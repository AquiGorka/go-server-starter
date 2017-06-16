package server_test

import (
	"context"
	"github.com/AquiGorka/go-server-starter/server"
	"github.com/go-speedo/go-speedo"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// run the server
	app := iris.New()
	app = server.HTTPServer(app)
	app = server.WebsocketServer(app)
	go app.Run(iris.Addr(":" + os.Getenv("APP_PORT")))
	// tests
	code := m.Run()
	// stop the server
	app.Shutdown(context.Background())
	// fin
	os.Exit(code)
}

func TestSocketServer(t *testing.T) {
	_, err := websocket.Dial("ws://localhost:"+os.Getenv("APP_PORT")+"/ws", "", "http://localhost/")
	if err != nil {
		t.Error("Error connecting to socket server: ", err)
	}
}

func TestHttpServer(t *testing.T) {
	_, err := http.Get("http://localhost:" + os.Getenv("APP_PORT") + "/ping")
	if err != nil {
		t.Error("Error connecting to http server: ", err)
	}
}
