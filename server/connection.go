package server

import (
	"github.com/go-speedo/go-speedo/websocket"
)

func onConnectionHandler(c websocket.Connection) {
	c.On("ping", func(msg string) { c.Emit("pong", "") })
}
