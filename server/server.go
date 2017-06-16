package server

import (
	"github.com/go-speedo/go-speedo"
	"github.com/go-speedo/go-speedo/context"
	"github.com/go-speedo/go-speedo/middleware/logger"
	"github.com/go-speedo/go-speedo/websocket"
)

// HTTPServer enhances an iris app to accept http requests
func HTTPServer(app *iris.Application) *iris.Application {
	// request logger
	customLogger := logger.New(logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
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

// WebsocketServer enhances an iris app to accept websocket connections
func WebsocketServer(app *iris.Application) *iris.Application {
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
