package server_test

import(
  "os"
  "testing"
  "golang.org/x/net/websocket"
  gws "github.com/gorilla/websocket"
  "github.com/rgamba/evtwebsocket"
)


func TestOnConnectionHandler(t *testing.T) {
  wsEndpoint := "ws://localhost:" + os.Getenv("APP_PORT") + "/ws"
  pingMsg := "iris-websocket-message:ping;0;all;"
  pongMsg := "iris-websocket-message:pong;0;"

  // net/websocket client
  ws, err := websocket.Dial(wsEndpoint, "", "http://localhost")
  if (err != nil) {
    t.Error("Client error connecting to websocket server: ", err)
  }
  buffer := []byte(pingMsg)
  _, err = ws.Write(buffer)
  if (err != nil) {
    t.Error("Client error could not write to socket: ", err)
  }
  var msg = make([]byte, 2048)
  var n int
  n, err = ws.Read(msg)
  if (err != nil) {
    t.Error("Client error reading from socket: ", err)
  }
  if (string(msg[:n]) != pongMsg) {
    t.Error("Client error did not receive pong event: ", string(msg[:n]))
  }

  // gorrilla client 
  var conn *gws.Conn
  conn, _, err = gws.DefaultDialer.Dial(wsEndpoint, nil)
  if (err != nil) {
    t.Error("Client error connecting to socket: ", err)
  }
  err = conn.WriteMessage(gws.TextMessage, []byte(pingMsg))
  if (err != nil) {
    t.Error("Client error writing message to socket: ", err)
  }
  _, msg, _ = conn.ReadMessage()
  if (string(msg) != pongMsg) {
    t.Error("Client error did not receive pong event: ", string(msg))
  }

  // evtwebsocket client
  // this waits for the async message listener to receive the pong message
  done := make(chan bool)
  go func() {
    c := evtwebsocket.Conn{
      OnMessage: func(msg []byte, w *evtwebsocket.Conn) {
        if (string(msg) != pongMsg) {
          t.Error("Client error did not receive pong event: ", string(msg))
        }
        done <- true
      },
      OnError: func(err error) {
        t.Error("Client error: ", err)
      },
    }
    err = c.Dial(wsEndpoint, "")
    if err != nil {
      t.Error("Client error connectiong to socket server: ", err)
    }
    c.Send(evtwebsocket.Msg{
      Body : []byte(pingMsg),
    })
  }()
  <-done
}

