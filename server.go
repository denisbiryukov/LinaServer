package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Server describes the receiver response mechanism.
type Server struct {
	isRunning bool           // The server status.
	listener  net.Listener   // The listening functionality.
	wg        sync.WaitGroup // The internal threads' synchronizer.
}

// RegisterSignals constructs a list of termination signals with further awaiting of any of them.
func (s *Server) RegisterSignals() {
	// Form a list of signals and a channel for them.
	signals := []os.Signal{os.Interrupt, os.Kill, syscall.SIGTERM}
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, signals...)

	// Block until a signal is received.
	signal := <-channel
	fmt.Println("Signal received: ", signal)

	fmt.Println("Graceful shutdown started.")
	s.isRunning = false
	s.listener.Close()
}

// Run starts the server process.
func (s *Server) Run() {
	// Register termination signals in a single thread.
	go s.RegisterSignals()

	// Open TCP port to listen.
	var err error = nil
	s.listener, err = net.Listen("tcp", fmt.Sprintf(":%v", Port))
	if err != nil {
		log.Fatalf("net.Listen() error: %v", err)
	}

	// Listen the port for incoming messages.
	for s.isRunning {
		connection, err := s.listener.Accept()
		if err != nil {
			if s.isRunning {
				log.Printf("listener.Accept() error: %v", err)
			}
			continue
		}
		// Handle connections concurrently.
		s.wg.Add(1)
		go s.Response(connection)
	}
	s.wg.Wait()
}

// Response handles an incoming connection and sends a message back to a client.
func (s *Server) Response(connection net.Conn) {
	defer s.wg.Done()
	defer connection.Close()
	log.Printf("Connection with %v", connection.RemoteAddr().String())

	seconds := GetSeconds(time.Now())
	// Process the outcoming message.
	m := Message(uint32(seconds))

	// Send an initial message to the server.
	err := binary.Write(connection, binary.BigEndian, m)
	if err != nil {
		log.Fatal("binary.Write failed:", err)
	}
	fmt.Println("Sent: ", seconds)
}
