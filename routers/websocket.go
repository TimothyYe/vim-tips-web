package routers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var connections map[*websocket.Conn]bool

func InitConnections() {
	connections = make(map[*websocket.Conn]bool)
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handleshake", 400)
		return
	} else if err != nil {
		fmt.Println(err)
	}

	connections[conn] = true
	fmt.Println("A new connection added...")
}

func sendAll(msg []byte) {
	fmt.Println("Send data to all clients...")
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(connections, conn)
			return
		}
	}
}
