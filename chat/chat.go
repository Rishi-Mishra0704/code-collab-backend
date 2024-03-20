package chat

import (
	"fmt"
	"time"

	"github.com/Rishi-Mishra0704/code-collab-backend/network"
)

// ChatService represents the chat service responsible for managing the chat system.
// It provides methods for sending and receiving messages, as well as managing connections with peers.
type ChatService struct {
	TCPTransport *network.TCPTransport // Reference to the TCPTransport instance
}

// NewChatService creates a new instance of ChatService with the provided transport layer.
// It initializes the chat service with the specified transport for communication.
func NewChatService(transport *network.TCPTransport) *ChatService {
	return &ChatService{
		TCPTransport: transport,
	}
}

// SendMessage sends a chat message to a specific room.
func (cs *ChatService) SendMessage(roomID string, sender *network.Peer, content string) error {
	// Create a new chat message with current timestamp
	timestamp := time.Now().Format("02-01-2006")

	// Retrieve the room from the TCPTransport
	cs.TCPTransport.Mutex.Lock()
	room, ok := cs.TCPTransport.Rooms[roomID]
	cs.TCPTransport.Mutex.Unlock()
	if !ok {
		return fmt.Errorf("room %s does not exist", roomID)
	}

	// Add the message to the chat history of the room
	room.Chat = append(room.Chat, fmt.Sprintf("[%s] %s: %s", timestamp, sender.ID, content))

	return nil
}
