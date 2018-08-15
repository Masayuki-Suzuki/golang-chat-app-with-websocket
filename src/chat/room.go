package main

type room struct {
	forward chan []byte      // a channel that holds messages for sending to another client.
	join    chan *client     // a channel for a client that will join chat room.
	leave   chan *client     // a channel for a client that will leave chat room.
	clients map[*client]bool // hold clients that are joining the chat room.
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			for client := range r.clients {
				select {
				case client.send <- msg:
					// Sending message:
				default:
					// Failed sending message:
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
