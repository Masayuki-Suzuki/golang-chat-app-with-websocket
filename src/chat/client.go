package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// "client" indicates one of the users who is chatting.
type client struct {
	socket   *websocket.Conn // socket is WebSocket for this client.
	send     chan *message   // send is channel sending messages.
	room     *room           // room is a chat room that this client is.
	userData map[string]interface{}
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
