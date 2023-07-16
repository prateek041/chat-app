package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	l *log.Logger
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func NewSocket(l *log.Logger) *WebSocket {
	return &WebSocket{l}
}

func (ws *WebSocket) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("this is the websocket endpoint")

	// upgrade the connection
	wsConnection, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		ws.l.Println("unable to upgrade the connection to Web Socket connection", err)
	}

	// pass it to the reader
	reader(ws, wsConnection)
}

func reader(ws *WebSocket, conn *websocket.Conn) {
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			ws.l.Println("Error in reading the message from connection", err)
		}

		// let's print the message sent from the frontend
		ws.l.Println("MESSAGE: ", string(msg))

		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}
