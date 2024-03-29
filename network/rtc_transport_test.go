package network

import (
	"testing"

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
