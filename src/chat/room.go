package main

type room struct {
	forward chan []byte      // a channel that holds messages for sending to another client.
	join    chan *client     // a channel for a client that will join chat room.
	leave   chan *client     // a channel for a client that will leave chat room.
	clients map[*client]bool // hold clients that are joining the chat room.
}
