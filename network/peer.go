package tcp

import "net"

// Peer represents a participant in the collaborative code editing network.
// Each peer is uniquely identified by an ID and may have associated metadata
// such as name, email, and network address. Peers are essential for
// facilitating communication, collaboration, and coordination within the
// distributed code editing environment.
type Peer struct {
	ID      string   // Unique identifier for the peer
	Name    string   // Name of the peer
	Email   string   // Email of the peer
	Address string   // Host:Port address of the peer
	Online  bool     // Indicates whether the peer is currently online
	Conn    net.Conn // Network connection for the peer
}
