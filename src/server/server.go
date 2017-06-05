package server

import (
  "gopkg.in/kataras/iris.v6"
  "gopkg.in/kataras/iris.v6/adaptors/httprouter"
  "gopkg.in/kataras/iris.v6/middleware/logger"
  "gopkg.in/kataras/iris.v6/adaptors/websocket"
)

func HttpServer(app *iris.Framework) *iris.Framework{
  // dev logger
  app.Adapt(iris.DevLogger())
  // router
  app.Adapt(httprouter.New())
  // request logger
  customLogger := logger.New(logger.Config{
    Status: true,
    IP: true,
    Method: true,
    Path: true,
  })
  app.Use(customLogger)
  // 404
  app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
    customLogger.Serve(ctx)
    notFoundHandler(ctx)
  })
  // ping
  app.Get("/ping", pingHandler)
  //
  return app
}

func WebsocketServer(app *iris.Framework) *iris.Framework{
  // endpoint
  ws := websocket.New(websocket.Config{
    Endpoint: "/ws",
  })
  // connection handler
  ws.OnConnection(onConnectionHandler)
  // websocket
  app.Adapt(ws)
  //
  return app
}
