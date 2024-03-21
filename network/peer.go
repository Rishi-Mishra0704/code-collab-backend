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
