package main

import (
	"encoding/binary"
	"log"
	"net"
)

// Client describes the sender request mechanism.
type Client struct {
}

// Request establishes a connection with a remote server. Sends a request to get a response.
func (c *Client) Request(host string) Message {
	// Establishes a connection with a remote server.
	connection, err := net.Dial("tcp", GetAddress(host))
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	// Process the incoming message.
	var message Message
	// Wait until the server responds.
	err = binary.Read(connection, binary.BigEndian, &message)
	if err != nil {
		log.Fatal("binary.Read failed:", err)
	}
	return message
}
