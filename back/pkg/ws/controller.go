package ws

import "log"

type Controller struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
}

func NewController() *Controller {
	return &Controller{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
	}
}

func (c *Controller) Run() {
	for {
		select {
		case client := <-c.Register:
			for cln := range c.Clients {
				cln.Conn.WriteJSON(Message{User: client.ID, Body: "New User Joined"})
			}
			c.Clients[client] = true
			client.Conn.WriteJSON(Message{User: client.ID, Body: "Me in"})
			log.Printf("Client %s has registered", client.ID)
		case client := <-c.Unregister:
			delete(c.Clients, client)
			for cln := range c.Clients {
				cln.Conn.WriteJSON(Message{User: client.ID, Body: "User Disconnected"})
			}
			client.Conn.WriteJSON(Message{User: client.ID, Body: "Me out"})
			log.Printf("Client %s has unregistered", client.ID)
		case message := <-c.Broadcast:
			log.Println("Sending message to all clients")
			for cln := range c.Clients {
				if err := cln.Conn.WriteJSON(message); err != nil {
					log.Printf("Controller: failed to broadcast to %s: %v", cln.ID, err)
				}
			}
		}
	}
}
