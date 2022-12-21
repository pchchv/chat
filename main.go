package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte // Message waiting to be sent.
}

// Keeps track of all connected clients,
// clients who are trying to register,
// clients who have been destroyed and are waiting to be deleted,
// and messages to and from all connected clients.
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var manager = ClientManager{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

func main() {
	fmt.Println("Starting application...")
	go manager.start()
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":12345", nil)
}
