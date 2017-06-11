package server

import (
  "github.com/kataras/iris"
  "github.com/kataras/iris/context"
)

func pingHandler(ctx context.Context) {
  ctx.StatusCode(iris.StatusOK)
  ctx.HTML("pong")
}

