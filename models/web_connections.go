package models

import (
	"github.com/gofiber/websocket/v2"
)

type Connections struct {
	clients []*websocket.Conn
}

func (c *Connections) AddConnection(conn *websocket.Conn) {
	c.clients = append(c.clients, conn)
}

func (c *Connections) RemoveConnection(conn *websocket.Conn) {
	for i, client := range c.clients {
		if client == conn {
			c.clients = append(c.clients[:i], c.clients[i+1:]...)
			break
		}
	}
}
