package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"trace"
)

type room struct {
	forward chan []byte      // a channel that holds messages for sending to another client.
	join    chan *client     // a channel for a client that will join chat room.
	leave   chan *client     // a channel for a client that will leave chat room.
	clients map[*client]bool // hold clients that are joining the chat room.
	tracer  trace.Tracer
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("Joined new client.")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Leave a client from the room.")
		case msg := <-r.forward:
			r.tracer.Trace("you got a new message.")
			for client := range r.clients {
				select {
				case client.send <- msg:
					// Sending message:
					r.tracer.Trace(" -- Your message has sent to the client.")
				default:
					// Failed sending message:
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" -- Failed sending message. System will clean up client.")
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
