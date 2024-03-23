package network

// Signal represents a signaling message used for establishing WebRTC connections between peers.
type Signal struct {
	Type string      // Type of signaling message (e.g., "offer", "answer", "ice")
	Data interface{} // Signaling data associated with the message
}

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

	// JoinRoom joins a collaborative editing room using the specified room ID.
	// It takes the ID of the room to join and performs any necessary actions to join the room.
	// Returns an error if joining the room fails.
	JoinRoom(roomID string, peer *Peer) error

	// LeaveRoom leaves the current collaborative editing room.
	// It leaves the current room and performs any necessary cleanup actions.
	// Returns an error if leaving the room fails.
	LeaveRoom(string, string) error

	// CreateRoom creates a new collaborative editing room and returns the room ID.
	// It creates a new room for collaborative editing and returns the ID of the newly created room.
	// Returns the room ID and an error if creating the room fails.
	CreateRoom(host *Peer) (string, error)

	// ExchangeSignal is used for exchanging signaling messages required for establishing
	// WebRTC connections between peers.
	ExchangeSignal(roomID string, peerID string, signal Signal) error

	// SendData sends data over the network to the specified peer in the specified room.
	// It sends data using the established WebRTC data channel between peers.
	SendData(roomID string, peerID string, data []byte) error

	// ReceiveData receives data over the network from the specified peer in the specified room.
	// It receives data from the WebRTC data channel between peers.
	ReceiveData(roomID string, peerID string) ([]byte, error)
}
