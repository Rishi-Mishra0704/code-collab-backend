package chat

import "github.com/Rishi-Mishra0704/code-collab-backend/network"

// ChatMessage represents a message in the chat system.
type ChatMessage struct {
	network.Peer
	Content   string
	Timestamp string
}
type ChatService struct {
	network.Transport // Embed the network.Transport interface
}

// NewChatService creates a new instance of ChatService.
func NewChatService(transport network.Transport) *ChatService {
	return &ChatService{
		Transport: transport,
	}
}
