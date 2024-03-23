package network

import "net"

// Peer represents a participant in the collaborative code editing network.
// Each peer is uniquely identified by an ID and may have associated metadata
// such as name, email, and network address. Peers are essential for
// facilitating communication, collaboration, and coordination within the
// distributed code editing environment.
type Peer struct {
	ID      string   `json:"id"`      // Unique identifier for the peer
	Name    string   `json:"name"`    // Name of the peer
	Email   string   `json:"email"`   // Email of the peer
	Address string   `json:"address"` // Host:Port address of the peer
	Online  bool     `json:"online"`  // Indicates whether the peer is currently online
	Conn    net.Conn `json:"-"`       // Network connection for the peer (omitted from JSON)
}

// Room represents a collaborative editing room in the network.
// It contains information about the room ID, host, connected peers, and chat history.
type Room struct {
	ID    string           `json:"id"`    // Unique identifier for the room
	Host  *Peer            `json:"host"`  // Peer representing the host of the room
	Peers map[string]*Peer `json:"peers"` // Map of connected peers in the room, keyed by peer ID
	Chat  []string         `json:"chat"`  // Chat history within the room
}
