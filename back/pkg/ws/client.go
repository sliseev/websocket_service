package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID         string
	Conn       *websocket.Conn
	Controller *Controller
}

type Message struct {
	User string `json:"user"`
	Body string `json:"body"`
}

func (c *Client) ReadName() error {
	_, p, err := c.Conn.ReadMessage()
	if err != nil {
		c.Controller.Unregister <- c
		c.Conn.Close()
		return err
	}
	c.ID = string(p)
	return nil
}

func (c *Client) ProcessMessages() {
	defer func() {
		c.Controller.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%s] Message Received: %s\n", c.ID, p)

		message := Message{User: c.ID, Body: string(p)}
		c.Controller.Broadcast <- message
	}
}
