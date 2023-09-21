package chatserver

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn
	In   chan Massage
	Err  chan error
}

type Massage struct {
	Sender, Reciever *Client
	Content          string
}
type Hub struct {
	In, Out chan Massage
	Control chan Client
}
