package chat

import "github.com/Rishi-Mishra0704/code-collab-backend/network"

// ChatMessage represents a message in the chat system.
// It encapsulates the content of the message, along with metadata such as the sender's information and the timestamp.
type ChatMessage struct {
	network.Peer        // Embed the network.Peer struct to represent both sender and receiver.
	Content      string // Content of the message.
	Timestamp    string // Timestamp indicating when the message was sent.
}

// ChatService represents the chat service responsible for managing the chat system.
// It provides methods for sending and receiving messages, as well as managing connections with peers.
type ChatService struct {
	network.Transport // Embed the network.Transport interface to handle communication with peers.
}

// NewChatService creates a new instance of ChatService with the provided transport layer.
// It initializes the chat service with the specified transport for communication.
func NewChatService(transport network.Transport) *ChatService {
	return &ChatService{
		Transport: transport,
	}
}
