package chat

import (
	"encoding/json"
	"fmt"
)

// serializeMessage serializes a ChatMessage into a byte slice.
func serializeMessage(msg *ChatMessage) ([]byte, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize message: %v", err)
	}
	return data, nil
}

// deserializeMessage deserializes a byte slice into a ChatMessage.
func deserializeMessage(data []byte) (*ChatMessage, error) {
	var msg ChatMessage
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize message: %v", err)
	}
	return &msg, nil
}
