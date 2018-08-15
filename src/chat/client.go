package main

import (
	"github.com/gorilla/websocket"
)

// "client" indicates one of the users who is chatting.
type client struct {
	socket *websocket.Conn // socket is WebSocket for this client.
	send   chan []byte     // send is channel sending messages.
	room   *room           // room is a chat room that this client is.
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}
