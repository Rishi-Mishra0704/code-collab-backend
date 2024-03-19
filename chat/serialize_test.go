package chat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeMessage_NoError(t *testing.T) {
	// Create a mock message with valid content
	msg := &ChatMessage{
		Content:   "Hello, world!",
		Timestamp: "2022-03-19",
	}

	// Serialize the message
	data, err := serializeMessage(msg)
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestDeserializeMessage_DeserializeError(t *testing.T) {
	// Create a mock byte slice with invalid JSON data that cannot be deserialized
	invalidData := []byte("invalid JSON data")

	// Attempt to deserialize the message
	msg, err := deserializeMessage(invalidData)

	// Ensure that an error is returned
	assert.Error(t, err)
	assert.Nil(t, msg)
}

func TestDeserializeMessage_NoError(t *testing.T) {
	// Create a mock message with valid content
	validData := []byte(`{"Content":"Hello, world!","Timestamp":"2022-03-19"}`)

	// Deserialize the message
	decodedMsg, err := deserializeMessage(validData)
	assert.NoError(t, err)
	assert.NotNil(t, decodedMsg)

	// Ensure that the deserialized message matches the original message
	assert.Equal(t, "Hello, world!", decodedMsg.Content)
	assert.Equal(t, "2022-03-19", decodedMsg.Timestamp)
}
