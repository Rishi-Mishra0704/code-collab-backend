package network

import (
	"net"
	"sync"
)

// TCPTransport implements the Transport interface using TCP.
type TCPTransport struct {
	// Listener is used to accept incoming connections
	Listener net.Listener

	// Peers is a map to store connected peers
	Peers map[string]*Peer
	// Mutex for safe access to the peers map
	Mutex sync.Mutex
}

// NewTCPTransport creates a new instance of TCPTransport.
func NewTCPTransport() *TCPTransport {
	return &TCPTransport{
		Peers: make(map[string]*Peer),
	}
}

// Listen starts listening for incoming TCP connections on the specified address.
func (t *TCPTransport) Listen(address string) error {
	// Check if listener is already initialized
	if t.Listener != nil {
		return nil // Listener already started
	}

	// Start listening for incoming TCP connections
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	t.Listener = listener
	return nil
}

// Close closes the TCP transport, releasing any associated resources.
func (t *TCPTransport) Close() error {
	// Check if listener is initialized
	if t.Listener == nil {
		return nil // Listener already closed
	}

	// Close the listener
	err := t.Listener.Close()
	if err != nil {
		return err
	}
	t.Listener = nil // Reset the listener
	return nil
}
