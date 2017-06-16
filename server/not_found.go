package server

import (
	"github.com/go-speedo/go-speedo"
	"github.com/go-speedo/go-speedo/context"
)

func notFoundHandler(ctx context.Context) {
	ctx.StatusCode(iris.StatusNotFound)
	ctx.HTML("404")
}
