package hub

import (
	"errors"

	"github.com/mamad-nik/chatserver"
)

type Hub struct {
	chatserver.Hub
	Clients map[*chatserver.Client]bool
}

func New() Hub {
	return Hub{
		Hub: chatserver.Hub{
			In:      make(chan chatserver.Massage),
			Out:     make(chan chatserver.Massage),
			Control: make(chan chatserver.Client),
		},
		Clients: make(map[*chatserver.Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Control:
			if _, ok := h.Clients[&client]; ok {
				delete(h.Clients, &client)
			} else {
				h.Clients[&client] = true
			}
		case massage := <-h.In:
			if _, ok := h.Clients[massage.Reciever]; ok {
				massage.Reciever.In <- massage
			} else {
				massage.Sender.Err <- errors.New("reciever not available")
			}
		}
	}
}
