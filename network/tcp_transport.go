package network

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"sync"
)

// TCPTransport implements the Transport interface using TCP.
// It manages the network transport layer responsible for facilitating communication between peers.
type TCPTransport struct {
	Listener net.Listener     // Listener for accepting incoming connections
	Rooms    map[string]*Room // Map to store rooms in the network, keyed by room ID
	Mutex    sync.Mutex       // Mutex for safe access to the rooms map
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

// CreateRoom creates a new collaborative editing room and returns the room ID.
// The room is created by taking the specified host as the initial peer in the room.
// It generates a unique room ID, creates a new room with the host, and adds the room to the network's rooms map.
// It returns the room ID and an error if creating the room fails.
func (t *TCPTransport) CreateRoom(host *Peer) (string, error) {
	roomID := generateRoomID() // Generate a unique room ID

	room := &Room{
		ID:    roomID,
		Host:  host,
		Peers: make(map[string]*Peer),
		Chat:  []string{},
	}

	if host.ID == "" || host.Name == "" || host.Address == "" || host.Email == "" {
		return "", errors.New("host peer must have ID, Name, Email and Address fields")
	}

	room.Peers[host.ID] = host // Add the host to the room

	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.Rooms[roomID] = room // Add the room to the network's rooms map

	return roomID, nil
}

// generateRoomID generates a random hexadecimal room ID.
// It generates 8 random bytes and converts them to a hexadecimal string.
func generateRoomID() string {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

// JoinRoom allows a peer to join a collaborative editing room by its ID.
// It checks if the specified room exists in the network.
// If the room exists, it adds the peer to the room's list of connected peers.
// It returns an error if the room doesn't exist or if the peer is already in the room.
func (t *TCPTransport) JoinRoom(roomID string, peer *Peer) error {
	t.Mutex.Lock()
	room, ok := t.Rooms[roomID]
	t.Mutex.Unlock()

	if !ok {
		return fmt.Errorf("room %s does not exist", roomID)
	}

	if _, exists := room.Peers[peer.ID]; exists {
		return fmt.Errorf("peer %s is already in room %s", peer.ID, roomID)
	}

	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	room.Peers[peer.ID] = peer // Add the peer to the room's connected peers

	fmt.Printf("Peer %s joined room %s\n", peer.ID, roomID)
	return nil
}

// LeaveRoom allows a peer to leave a collaborative editing room by its ID.
// It checks if the specified room exists in the network.
// If the room exists, it removes the peer from the room's list of connected peers.
// It returns an error if the room doesn't exist or if the peer is not in the room.
func (t *TCPTransport) LeaveRoom(roomID string, peerID string) error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	room, ok := t.Rooms[roomID]
	if !ok {
		return fmt.Errorf("room %s does not exist", roomID)
	}

	_, exists := room.Peers[peerID]
	if !exists {
		return fmt.Errorf("peer %s is not in room %s", peerID, roomID)
	}

	delete(room.Peers, peerID)
	fmt.Printf("Peer %s left room %s\n", peerID, roomID)

	return nil
}

func (t *TCPTransport) GetAllRooms() map[string]*Room {
	return t.Rooms
}

// ExchangeSignal is not required for TCP transport.
func (t *TCPTransport) ExchangeSignal(roomID string, peerID string, signal Signal) error {
	return nil
}

// SendData is not required for TCP transport.
// It would typically be used for sending data over WebRTC data channels.
func (t *TCPTransport) SendData(roomID string, peerID string, data []byte) error {
	return nil
}

// ReceiveData is not required for TCP transport.
// It would typically be used for receiving data over WebRTC data channels.
func (t *TCPTransport) ReceiveData(roomID string, peerID string) ([]byte, error) {
	return nil, nil
}
