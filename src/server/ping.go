package server

import (
  "gopkg.in/kataras/iris.v6"
)

func pingHandler(ctx *iris.Context) {
  ctx.HTML(iris.StatusOK, "pong")
}

