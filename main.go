package main

import(
  "os"
  "github.com/AquiGorka/go-server-starter/server"
  "github.com/kataras/iris"
)

func main() {
  app := iris.New()
  // http server
  app = server.HttpServer(app)
  // websocket server
  app = server.WebsocketServer(app)
  //
  app.Run(iris.Addr(":" + os.Getenv("APP_PORT")))
}
