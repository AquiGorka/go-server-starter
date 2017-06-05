package server

import(
  "gopkg.in/kataras/iris.v6/adaptors/websocket"
)

func onConnectionHandler(c websocket.Connection) {
  c.On("ping", func(msg string) { c.Emit("pong", "") })
}
