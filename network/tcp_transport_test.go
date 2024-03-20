package network

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport_Listen(t *testing.T) {
	transport := NewTCPTransport()

	// Test successful listener initialization
	address := SetupTest(t) // Use random port
	err := transport.Listen(address)
	assert.NoError(t, err)
	assert.NotNil(t, transport.Listener)

	// Test attempting to listen on the same address again
	err = transport.Listen(address)
	assert.NoError(t, err) // No error should occur since listener is already started
}

func TestTCPTransport_Listen_Error(t *testing.T) {
	transport := NewTCPTransport()

	// Test listener initialization with an invalid address
	err := transport.Listen("invalid_address")
	assert.Error(t, err)
	assert.Nil(t, transport.Listener)
}

func TestTCPTransport_Close(t *testing.T) {
	transport := NewTCPTransport()
	addr := SetupTest(t)
	// Initialize listener
	err := transport.Listen(addr) // Use random port
	assert.NoError(t, err)

	// Test closing the listener
	err = transport.Close()
	assert.NoError(t, err)
	assert.Nil(t, transport.Listener)
}

func TestTCPTransport_Close_NotInitialized(t *testing.T) {
	transport := NewTCPTransport()

	// Test closing when listener is not initialized
	err := transport.Close()
	assert.NoError(t, err) // No error should occur if listener is not initialized
}

func TestTCPTransport_CreateRoom(t *testing.T) {
	transport := NewTCPTransport()
	addr := SetupTest(t)
	host := &Peer{
		ID:      "host_peer_id",
		Name:    "Host Peer",
		Email:   "host@example.com",
		Address: addr,
		Online:  true,
		Conn:    nil, // Connection not needed for this test
	}

	// Create a room
	roomID, err := transport.CreateRoom(host)
	assert.NoError(t, err)

	// Check that the room is added to the map of rooms
	transport.Mutex.Lock()
	defer transport.Mutex.Unlock()
	assert.Contains(t, transport.Rooms, roomID)

	// Check that the room contains the host
	room := transport.Rooms[roomID]
	assert.NotNil(t, room)
	assert.Equal(t, host, room.Host)
	assert.Contains(t, room.Peers, host.ID)
}

// TestJoinRoom tests the JoinRoom function of the TCPTransport.
func TestJoinRoom(t *testing.T) {
	// Create a new TCPTransport instance
	transport := NewTCPTransport()

	// Create a peer to act as the host
	host := &Peer{
		ID: "host1",
	}

	// Create a room and add the host
	roomID, err := transport.CreateRoom(host)
	if err != nil {
		t.Fatalf("Failed to create room: %v", err)
	}

	// Create a peer to join the room
	peer := &Peer{
		ID: "peer1",
	}

	// Join the room with the peer
	err = transport.JoinRoom(roomID, peer)
	if err != nil {
		t.Fatalf("Failed to join room: %v", err)
	}

	// Attempt to join the room again with the same peer
	err = transport.JoinRoom(roomID, peer)
	if err == nil {
		t.Fatalf("JoinRoom did not return error for duplicate peer")
	} else if err.Error() != fmt.Sprintf("peer %s is already in room %s", peer.ID, roomID) {
		t.Fatalf("JoinRoom returned unexpected error message for duplicate peer")
	}

	// Attempt to join a non-existent room
	err = transport.JoinRoom("nonexistent", peer)
	if err == nil {
		t.Fatalf("JoinRoom did not return error for non-existent room")
	} else if err.Error() != fmt.Sprintf("room %s does not exist", "nonexistent") {
		t.Fatalf("JoinRoom returned unexpected error message for non-existent room")
	}
}

// TestGenerateRoomID tests the generateRoomID function.
func TestGenerateRoomID(t *testing.T) {
	// Generate a room ID
	roomID := generateRoomID()

	// Check if the room ID is of the correct length (16 characters for 8 bytes)
	if len(roomID) != 16 {
		t.Fatalf("Generated room ID has incorrect length: got %d, want 16", len(roomID))
	}

	// Ensure that the room ID consists of valid hexadecimal characters
	_, err := hex.DecodeString(roomID)
	if err != nil {
		t.Fatalf("Generated room ID contains invalid hexadecimal characters: %s", err)
	}
}

// TestLeaveRoom tests the LeaveRoom function of the TCPTransport.
func TestLeaveRoom(t *testing.T) {
	// Create a new TCPTransport instance
	transport := NewTCPTransport()

	// Create a peer to act as the host
	host := &Peer{
		ID: "host1",
	}

	// Create a room and add the host
	roomID, err := transport.CreateRoom(host)
	if err != nil {
		t.Fatalf("Failed to create room: %v", err)
	}

	// Create a peer to join the room
	peer := &Peer{
		ID: "peer1",
	}
	err = transport.JoinRoom(roomID, peer)
	if err != nil {
		t.Fatalf("Failed to join room: %v", err)
	}

	// Leave the room with the peer
	err = transport.LeaveRoom(roomID, peer.ID)
	if err != nil {
		t.Fatalf("Failed to leave room: %v", err)
	}

	// Attempt to leave the room again with the same peer
	err = transport.LeaveRoom(roomID, peer.ID)
	if err == nil {
		t.Fatalf("LeaveRoom did not return error for peer not in room")
	}

	// Attempt to leave a non-existent room
	err = transport.LeaveRoom("nonexistent", peer.ID)
	if err == nil {
		t.Fatalf("LeaveRoom did not return error for non-existent room")
	}
}
