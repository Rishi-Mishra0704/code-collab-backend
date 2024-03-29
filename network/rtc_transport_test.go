package network

import (
	"fmt"
	"testing"

	"github.com/pion/webrtc/v3"
	"github.com/stretchr/testify/assert"
)

// MockPeer represents a mock implementation of the Peer struct for testing purposes.
type MockPeer struct {
	*Peer // Include a field of type *Peer
}

func TestCreateRoom(t *testing.T) {
	// Initialize RTCTransport
	rt := &RTCTransport{
		Rooms: make(map[string]*RTCRoom),
	}

	// Mock peer
	mockPeer := &MockPeer{
		Peer: &Peer{
			ID:      "mock_peer_id",
			Name:    "Mock Peer",
			Email:   "mock@example.com",
			Address: "localhost:1234",
			Online:  true,
			Conn:    nil,
		},
	}

	// Call CreateRoom
	roomID, err := rt.CreateRoom(mockPeer.Peer)

	// Assert whether the room was created successfully
	assert.NoError(t, err)
	assert.NotNil(t, rt.Rooms[roomID])
}

func TestJoinRoom(t *testing.T) {
	// Create a new instance of RTCTransport for testing
	transport := &RTCTransport{
		Rooms: make(map[string]*RTCRoom),
	}

	// Create a room ID
	roomID := "room1"

	// Create a host peer
	host := &Peer{
		ID:      "host1",
		Name:    "Host Peer",
		Email:   "host@example.com",
		Address: "localhost:1234",
		Online:  true,
		Conn:    nil,
	}

	// Create the room and add the host
	transport.Rooms[roomID] = &RTCRoom{
		Room:    &Room{ID: roomID, Host: host, Peers: make(map[string]*Peer)},
		HostPC:  nil, // Assuming not relevant for this test
		PeerPCs: make(map[string]*webrtc.PeerConnection),
	}

	// Create a peer to join the room
	peer := &Peer{
		ID:      "peer1",
		Name:    "Joining Peer",
		Email:   "joiner@example.com",
		Address: "localhost:5678", // Assuming different address from the host
		Online:  true,
		Conn:    nil,
	}

	// Test joining the room with the peer
	err := transport.JoinRoom(roomID, peer)

	// Assert whether the peer joined the room successfully
	assert.NoError(t, err, "JoinRoom should not return an error")

	// Attempt to join the room again with the same peer
	err = transport.JoinRoom(roomID, peer)

	// Assert that it returns an error for duplicate peer
	assert.Error(t, err, "JoinRoom should return an error for duplicate peer")
	assert.Contains(t, err.Error(), fmt.Sprintf("peer %s is already in room %s", peer.ID, roomID), "Error message should indicate peer is already in the room")

	// Attempt to join a non-existent room
	err = transport.JoinRoom("nonexistent", peer)

	// Assert that it returns an error for non-existent room
	assert.Error(t, err, "JoinRoom should return an error for non-existent room")
	assert.Contains(t, err.Error(), "room nonexistent does not exist", "Error message should indicate non-existent room")
}
