package chat

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Rishi-Mishra0704/code-collab-backend/network"
)

// SetupRandomAddr generates a random address for testing purposes.
func SetupRandomAddr(t *testing.T) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	port := random.Intn(1000) + 50000 // Random port between 50000 and 50999
	address := fmt.Sprintf("localhost:%d", port)

	return address
}

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
