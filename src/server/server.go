package server

import (
  "github.com/kataras/iris"
  "github.com/kataras/iris/websocket"
  "github.com/kataras/iris/middleware/logger"
  "github.com/kataras/iris/context"
)

func HttpServer(app *iris.Application) *iris.Application{
  // request logger
  customLogger := logger.New(logger.Config{
    Status: true,
    IP: true,
    Method: true,
    Path: true,
  })
  app.Use(customLogger)
  // 404
  app.OnErrorCode(iris.StatusNotFound, func(ctx context.Context) {
    customLogger(ctx)
    notFoundHandler(ctx)
  })
  // ping
  app.Get("/ping", pingHandler)
  //
  return app
}

func WebsocketServer(app *iris.Application) *iris.Application{
  // endpoint
  ws := websocket.New(websocket.Config{
    Endpoint: "/ws",
  })
  // connection handler
  ws.OnConnection(onConnectionHandler)
  // websocket
  ws.Attach(app)
  //
  return app
}
