package client

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mamad-nik/chatserver"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ClientHandler(hub *chatserver.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	client := chatserver.Client{Conn: conn}
	hub.Control <- client
}
