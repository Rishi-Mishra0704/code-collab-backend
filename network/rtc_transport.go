package network

import (
	"errors"
	"fmt"
	"sync"

	"github.com/pion/webrtc/v3"
)

// RTCRoom represents a collaborative editing room in the network using WebRTC.
type RTCRoom struct {
	Room    *Room                             // Original room information
	HostPC  *webrtc.PeerConnection            // PeerConnection instance for the room host
	PeerPCs map[string]*webrtc.PeerConnection // Map of PeerConnection instances for connected peers, keyed by peer ID
}

// RTCTransport represents the network transport layer for WebRTC communication.
type RTCTransport struct {
	Rooms map[string]*RTCRoom // Map of active rooms, keyed by room ID
	Mutex sync.Mutex
	// Additional fields specific to WebRTC transport layer (e.g., ICE servers, configurations, etc.)
}

var _ Transport = (*RTCTransport)(nil)

func (t *RTCTransport) JoinRoom(roomID string, peer *Peer) error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	room, ok := t.Rooms[roomID]
	if !ok {
		return fmt.Errorf("room %s does not exist", roomID)
	}

	// Check if the peer is already in the room
	if _, exists := room.PeerPCs[peer.ID]; exists {
		return fmt.Errorf("peer %s is already in room %s", peer.ID, roomID)
	}

	// Add the peer to the room's connected peers
	room.PeerPCs[peer.ID] = nil // Initialize to nil for now, assuming it will be initialized later

	fmt.Printf("Peer %s joined room %s\n", peer.ID, roomID)
	return nil
}

// LeaveRoom leaves the current collaborative editing room.
// It closes the WebRTC PeerConnection and removes the peer from the room.
func (rt *RTCTransport) LeaveRoom(roomID string, peerID string) error {
	// Implementation to close WebRTC PeerConnection and remove the peer from the room
	return nil
}

// CreateRoom creates a new collaborative editing room and returns the room ID.
// It initializes a new RTCRoom instance and sets up the host's WebRTC PeerConnection.
func (rt *RTCTransport) CreateRoom(host *Peer) (string, error) {
	// Generate a unique room ID
	roomID := generateRoomID() // You need to implement a function to generate a unique room ID

	// Initialize a new RTCRoom instance
	room := &RTCRoom{
		Room: &Room{
			ID: roomID,
			// You can set other room information here if needed
		},
		PeerPCs: make(map[string]*webrtc.PeerConnection),
	}

	// Initialize the host's WebRTC PeerConnection
	hostPC, err := initializePeerConnection()
	if err != nil {
		return "", err
	}

	// Set the host's WebRTC PeerConnection in the RTCRoom
	room.HostPC = hostPC

	// Add the room to the Rooms map
	rt.Rooms[roomID] = room

	return roomID, nil
}

// ExchangeSignal is used for exchanging signaling messages required for establishing
// WebRTC connections between peers.
func (rt *RTCTransport) ExchangeSignal(roomID string, peerID string, signal Signal) error {
	// Implementation to exchange signaling messages between peers using WebRTC
	return nil
}

// SendData sends data over the network to the specified peer in the specified room.
// It sends data using the established WebRTC data channel between peers.
func (rt *RTCTransport) SendData(roomID string, peerID string, data []byte) error {
	// Implementation to send data over WebRTC data channel
	return nil
}

// ReceiveData receives data over the network from the specified peer in the specified room.
// It receives data from the WebRTC data channel between peers.
func (rt *RTCTransport) ReceiveData(roomID string, peerID string) ([]byte, error) {
	// Implementation to receive data from WebRTC data channel
	return nil, nil
}
func initializePeerConnection() (*webrtc.PeerConnection, error) {
	// Implement your logic to initialize a new WebRTC PeerConnection
	// Example:
	config := webrtc.Configuration{}
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return nil, errors.New("failed to create PeerConnection: " + err.Error())
	}
	return peerConnection, nil
}

// Close method is not required for WebRTC PeerConnection
func (t *RTCTransport) Close() error {
	return nil
}

func (t *RTCTransport) Listen(address string) error {
	return nil
}
