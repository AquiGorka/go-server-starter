package server

import (
  "gopkg.in/kataras/iris.v6"
)

func notFoundHandler(ctx *iris.Context) {
  ctx.HTML(iris.StatusNotFound, "404")
}
