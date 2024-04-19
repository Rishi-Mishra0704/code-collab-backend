package network

import (
	"net"
	"sync"
)

// TCPTransport implements the Transport interface using TCP.
// It manages the network transport layer responsible for facilitating communication between peers.
type TCPTransport struct {
	Listener net.Listener     // Listener for accepting incoming connections
	Mutex    sync.Mutex       // Mutex for safe access to the rooms map
	Rooms    map[string]*Room // Map to store rooms in the network, keyed by room ID
}

var _ Transport = (*TCPTransport)(nil)

// NewTCPTransport creates a new instance of TCPTransport.
// It initializes the Rooms map to store rooms in the network.
func NewTCPTransport() *TCPTransport {
	return &TCPTransport{
		Rooms: make(map[string]*Room),
	}
}

// Listen starts listening for incoming TCP connections on the specified address.
// It initializes the network listener if not already initialized.
func (t *TCPTransport) Listen(address string) error {
	if t.Listener != nil {
		return nil // Listener already started
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	t.Listener = listener
	return nil
}

// Close closes the TCP transport, releasing any associated resources.
// It closes the network listener if it's initialized.
func (t *TCPTransport) Close() error {
	if t.Listener == nil {
		return nil // Listener already closed
	}

	err := t.Listener.Close()
	if err != nil {
		return err
	}
	t.Listener = nil // Reset the listener
	return nil
}
