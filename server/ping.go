package server

import (
	"github.com/go-speedo/go-speedo"
	"github.com/go-speedo/go-speedo/context"
)

func pingHandler(ctx context.Context) {
	ctx.StatusCode(iris.StatusOK)
	ctx.HTML("pong")
}
