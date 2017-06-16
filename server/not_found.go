package server

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func notFoundHandler(ctx context.Context) {
	ctx.StatusCode(iris.StatusNotFound)
	ctx.HTML("404")
}
