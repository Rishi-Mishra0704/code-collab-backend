package chat

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Rishi-Mishra0704/code-collab-backend/network"
)

func TestSendMessage(t *testing.T) {
	// Create a new instance of TCPTransport
	transport := &network.TCPTransport{
		Rooms: make(map[string]*network.Room),
	}

	// Create a new instance of ChatService
	chatService := NewChatService(transport)
	addr := SetupRandomAddr(t)

	// Create a test peer
	sender := &network.Peer{
		ID:      "sender1",
		Address: addr,
	}

	// Create a test room
	roomID := "room1"
	room := &network.Room{
		ID:    roomID,
		Host:  sender,
		Peers: make(map[string]*network.Peer),
		Chat:  []string{},
	}
	transport.Rooms[roomID] = room

	// Send a message
	content := "Hello, world!"
	err := chatService.SendMessage(roomID, sender, content)

	// Check if there are any errors
	assert.NoError(t, err, "SendMessage should not return an error")

	// Check if the message is added to the chat history
	assert.Len(t, transport.Rooms[roomID].Chat, 1, "Chat history should contain one message")
	assert.Contains(t, transport.Rooms[roomID].Chat[0], content, "Chat history should contain the sent message")
}

// TestSendMessageRoomNotExist tests the SendMessage method when the specified room does not exist.
func TestSendMessageRoomNotExist(t *testing.T) {
	// Create a new instance of TCPTransport
	transport := &network.TCPTransport{
		Rooms: make(map[string]*network.Room),
	}

	// Create a new instance of ChatService
	chatService := NewChatService(transport)

	// Create a test peer
	sender := &network.Peer{
		ID: "sender1",
	}

	// Attempt to send a message to a non-existent room
	roomID := "nonexistent"
	content := "Hello, world!"
	err := chatService.SendMessage(roomID, sender, content)

	// Check if the error is returned as expected
	assert.Error(t, err, "SendMessage should return an error when the room does not exist")
	assert.EqualError(t, err, fmt.Sprintf("room %s does not exist", roomID), "Error message should indicate that the room does not exist")
}
