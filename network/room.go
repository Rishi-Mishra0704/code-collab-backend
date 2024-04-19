package network

// HandleRoom is an interface that defines the methods for handling collaborative editing rooms.
type HandleRoom interface {

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
}
