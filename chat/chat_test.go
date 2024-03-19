package chat

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Rishi-Mishra0704/code-collab-backend/network"
)

func TestChatService_SendMessage(t *testing.T) {
	// Mock transport layer
	mockTransport := &MockTransport{
		SendFunc: func(data []byte, peer *network.Peer) error {
			// Verify that the message data is not empty
			assert.NotEmpty(t, data)

			// Verify that the peer is not nil
			assert.NotNil(t, peer)

			return nil
		},
	}

	// Create a new chat service with the mock transport
	chatService := NewChatService(mockTransport)
	// Create a mock address for the peer
	addr := SetupTest(t)
	// Create a mock peer
	mockPeer := &network.Peer{
		ID:      "mock_peer_id",
		Name:    "Mock Peer",
		Address: addr,
		Online:  true,
	}

	// Send a message
	err := chatService.SendMessage(mockPeer, "Hello, world!")
	assert.NoError(t, err)

	t.Run("SendError", func(t *testing.T) {
		// Set up the mock transport to return an error when sending the message
		mockTransport.SendFunc = func(data []byte, peer *network.Peer) error {
			return errors.New("send error")
		}

		// Send a message
		err := chatService.SendMessage(mockPeer, "Hello, world!")
		// Ensure that an error is returned
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to send message")
		assert.Contains(t, err.Error(), "send error")
	})

}
