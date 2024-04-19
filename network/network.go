package network

// Transport represents the network transport layer responsible for facilitating
// communication between peers in the collaborative code editing network.
// It supports TCP and WebRTC protocols but can be extended to support other protocols(UDP, Sockets, RPC, ...).
type Transport interface {

	// Listen starts listening for incoming connections on the specified address.
	// It takes the address (host:port) on which to listen for incoming connections.
	// Returns an error if starting the listener fails.
	Listen(address string) error

	// Close closes the network transport, releasing any associated resources.
	// It stops listening for incoming connections and cleans up any resources used by the transport.
	// Returns an error if closing the transport fails.
	Close() error
}
