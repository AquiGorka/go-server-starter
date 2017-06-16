package server

import (
	"github.com/kataras/iris/websocket"
)

func onConnectionHandler(c websocket.Connection) {
	c.On("ping", func(msg string) { c.Emit("pong", "") })
}
