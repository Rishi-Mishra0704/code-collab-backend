package network

import (
	"crypto/rand"
	"encoding/hex"
	"net"
	"sync"
)

type Room struct {
	ID    string
	Host  *Peer
	Peers map[string]*Peer
	Chat  []string
}

// TCPTransport implements the Transport interface using TCP.
type TCPTransport struct {
	// Listener is used to accept incoming connections
	Listener net.Listener

	// Rooms is a map to store rooms in a network
	Rooms map[string]*Room
	// Mutex for safe access to the peers map
	Mutex sync.Mutex
}

// NewTCPTransport creates a new instance of TCPTransport.
func NewTCPTransport() *TCPTransport {
	return &TCPTransport{
		Rooms: make(map[string]*Room),
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

func (t *TCPTransport) CreateRoom(host *Peer) (string, error) {
	// Generate a unique room ID
	roomID := generateRoomID()

	// Create a new room
	room := &Room{
		ID:    roomID,
		Host:  host,
		Peers: make(map[string]*Peer),
		Chat:  []string{},
	}

	// Add the host to the room
	room.Peers[host.ID] = host

	// Add the room to the map of connected rooms
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.Rooms[roomID] = room

	return roomID, nil
}

func generateRoomID() string {
	// Generate 8 random bytes
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		panic(err) // Error handling can be adjusted based on the use case
	}

	// Convert the random bytes to a hexadecimal string
	roomID := hex.EncodeToString(bytes)

	return roomID
}
