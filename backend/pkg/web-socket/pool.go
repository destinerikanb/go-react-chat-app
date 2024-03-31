package web_socket

import "fmt"

type Pool struct {
	// Channel to register new client
	Register chan *Client
	// Channel to unregister client
	Unregister chan *Client
	// Map of connected clients
	Clients map[*Client]bool
	// Channel to broadcast message to all of clients
	Broadcast chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println("Client:", client)
				if client.Conn != nil {
					err := client.Conn.WriteJSON(Message{
						Type: 1,
						Body: "New User Joined...",
					})

					if err != nil {
						fmt.Println("Error writing JSON:", err)
					}
				} else {
					fmt.Println("Client connection is nil")
				}
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{
					Type: 1,
					Body: "User Disconnected...",
				})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
