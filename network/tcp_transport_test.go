package network

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockReader struct{}

// Read always returns an error.
func (m *MockReader) Read(b []byte) (n int, err error) {
	return 0, errors.New("error")
}

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
		// Connection not needed for this test
	}

	// Create a room
	roomID, err := transport.CreateRoom(host)
	assert.NoError(t, err)
	assert.NotEmpty(t, roomID)

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
	addr := SetupTest(t)
	host := &Peer{
		ID:      "host1",
		Name:    "Host Peer",
		Email:   "host@example.com",
		Address: addr,
		Online:  true,
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

	if roomID == "" {
		t.Errorf("Error generating room ID")
	}
	// Ensure that the room ID consists of valid hexadecimal characters
	_, err := hex.DecodeString(roomID)
	if err != nil {
		t.Fatalf("Generated room ID contains invalid hexadecimal characters: %s", err)
	}
}

func TestGenerateRoomID_Error(t *testing.T) {
	// Replace rand.Reader with MockReader
	randReaderOrig := rand.Reader
	rand.Reader = &MockReader{}
	defer func() { rand.Reader = randReaderOrig }()

	// Generate a room ID
	roomID := generateRoomID()

	if roomID != "" {
		t.Errorf("Expected empty room ID when error occurs, got: %s", roomID)
	}
}

// TestLeaveRoom tests the LeaveRoom function of the TCPTransport.
func TestLeaveRoom(t *testing.T) {
	// Create a new TCPTransport instance
	transport := NewTCPTransport()

	// Create a peer to act as the host
	addr := SetupTest(t)
	host := &Peer{
		ID:      "host1",
		Name:    "Host Peer",
		Email:   "host@example.com",
		Address: addr,
		Online:  true,
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

func TestGetAllRooms(t *testing.T) {
	// Create a new instance of TCPTransport
	tcpTransport := NewTCPTransport()

	// Create some sample rooms
	room1 := &Room{
		ID:    "room1",
		Host:  &Peer{ID: "host1"},
		Peers: map[string]*Peer{},
		Chat:  []string{},
	}
	room2 := &Room{
		ID:    "room2",
		Host:  &Peer{ID: "host2"},
		Peers: map[string]*Peer{},
		Chat:  []string{},
	}

	// Add the rooms to the TCPTransport
	tcpTransport.Rooms["room1"] = room1
	tcpTransport.Rooms["room2"] = room2

	// Call the GetAllRooms method
	rooms := tcpTransport.GetAllRooms()

	// Check if the returned map of rooms is correct
	expectedRooms := map[string]*Room{
		"room1": room1,
		"room2": room2,
	}

	// Use assert.Equal to check if the actual and expected rooms are equal
	assert.Equal(t, expectedRooms, rooms, "GetAllRooms() returned incorrect rooms")
}

func TestCreateRoomError(t *testing.T) {
	transport := NewTCPTransport()
	host := &Peer{}
	roomID, err := transport.CreateRoom(host)
	assert.Error(t, err)
	assert.Empty(t, roomID)

}
