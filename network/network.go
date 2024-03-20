package network

// Transport represents the network transport layer responsible for facilitating
// communication between peers in the collaborative code editing network.
// It supports TCP and WebRTC protocols but can be extended to support other protocols(UDP, Sockets, RPC, ...).
type Transport interface {
	// Send sends data to the specified peer over the network.
	// It takes a byte slice containing the data to be sent and a pointer to the destination peer.
	// Returns an error if sending the data fails.
	Send(data []byte, peer *Peer) error

	// Receive receives data from any peer over the network.
	// It returns the received data as a byte slice and the source peer from which the data was received.
	// Returns an error if receiving data fails.
	Receive() ([]byte, *Peer, error)

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
	JoinRoom(roomID string) error

	// LeaveRoom leaves the current collaborative editing room.
	// It leaves the current room and performs any necessary cleanup actions.
	// Returns an error if leaving the room fails.
	LeaveRoom() error

	// CreateRoom creates a new collaborative editing room and returns the room ID.
	// It creates a new room for collaborative editing and returns the ID of the newly created room.
	// Returns the room ID and an error if creating the room fails.
	CreateRoom(host *Peer) (string, error)

	Connect(peer *Peer)
}
