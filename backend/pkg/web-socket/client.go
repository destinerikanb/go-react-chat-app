package web_socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Name string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	// closing websocket connection and delete client at the end
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	// Continuously read message from websocket connection
	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Send object Message to Pool and boradcast to all client
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
