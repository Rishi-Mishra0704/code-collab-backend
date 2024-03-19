package chat

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/Rishi-Mishra0704/code-collab-backend/network"
)

// SetupTest creates a random address for testing.
func SetupTest(t *testing.T) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	port := random.Intn(1000) + 50000 // Random port between 50000 and 50999
	address := fmt.Sprintf("localhost:%d", port)

	return address
}

// TestMain is the main entry point for tests.
func TestMain(m *testing.M) {
	// Add any global setup code here if needed
	exitVal := m.Run()
	// Add any global teardown code here if needed
	os.Exit(exitVal)
}

// MockTransport is a mock implementation of the network.Transport interface for testing.
type MockTransport struct {
	SendFunc       func(data []byte, peer *network.Peer) error
	ReceiveFn      func() ([]byte, *network.Peer, error)
	ListenFn       func(address string) error
	CloseFn        func() error
	ConnectFunc    func(address string) (*network.Peer, error) // Add Connect function
	JoinRoomFunc   func(roomID string) error                   // Add JoinRoom function
	LeaveRoomFunc  func() error                                // Add LeaveRoom function
	CreateRoomFunc func() (string, error)                      // Add CreateRoom function
}

// Send sends data to the specified peer.
func (m *MockTransport) Send(data []byte, peer *network.Peer) error {
	if m.SendFunc != nil {
		return m.SendFunc(data, peer)
	}
	return errors.New("SendFunc not implemented")
}

// Receive receives data from any peer.
func (m *MockTransport) Receive() ([]byte, *network.Peer, error) {
	if m.ReceiveFn != nil {
		return m.ReceiveFn()
	}
	return nil, nil, errors.New("ReceiveFn not implemented")
}

// Listen starts listening for incoming connections.
func (m *MockTransport) Listen(address string) error {
	if m.ListenFn != nil {
		return m.ListenFn(address)
	}
	return errors.New("ListenFn not implemented")
}

// Close closes the transport.
func (m *MockTransport) Close() error {
	if m.CloseFn != nil {
		return m.CloseFn()
	}
	return errors.New("CloseFn not implemented")
}

// Connect connects to a peer with the specified address.
func (m *MockTransport) Connect(address string) (*network.Peer, error) {
	if m.ConnectFunc != nil {
		return m.ConnectFunc(address)
	}
	return nil, errors.New("ConnectFunc not implemented")
}

// JoinRoom joins a collaborative editing room using the specified room ID.
func (m *MockTransport) JoinRoom(roomID string) error {
	if m.JoinRoomFunc != nil {
		return m.JoinRoomFunc(roomID)
	}
	return errors.New("JoinRoomFunc not implemented")
}

// LeaveRoom leaves the current collaborative editing room.
func (m *MockTransport) LeaveRoom() error {
	if m.LeaveRoomFunc != nil {
		return m.LeaveRoomFunc()
	}
	return errors.New("LeaveRoomFunc not implemented")
}

// CreateRoom creates a new collaborative editing room and returns the room ID.
func (m *MockTransport) CreateRoom() (string, error) {
	if m.CreateRoomFunc != nil {
		return m.CreateRoomFunc()
	}
	return "", errors.New("CreateRoomFunc not implemented")
}
