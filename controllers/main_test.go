package controllers

import (
	"github.com/Rishi-Mishra0704/code-collab-backend/network"
)

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

// func generateRoomID() string {
// 	bytes := make([]byte, 8)
// 	if _, err := rand.Read(bytes); err != nil {
// 		fmt.Printf("error creating room id : %s", err)
// 	}
// 	return hex.EncodeToString(bytes)
// }
