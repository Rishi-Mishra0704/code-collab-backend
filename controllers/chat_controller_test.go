package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rishi-Mishra0704/code-collab-backend/chat"
	"github.com/Rishi-Mishra0704/code-collab-backend/network"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateRoom(t *testing.T) {
	// Create a new instance of the TCPTransport
	transport := network.NewTCPTransport()
	chatService := chat.NewChatService(transport)

	// Create a new instance of the ChatController
	chatController := NewChatController(transport, chatService)

	// Create a new Gin router instance
	router := gin.Default()

	// Define the route for CreateRoom
	router.POST("/rooms", chatController.CreateRoom)

	// Test case 1: Successful room creation
	t.Run("SuccessfulRoomCreation", func(t *testing.T) {
		// Create a sample host peer
		host := &network.Peer{
			ID:      "host123",
			Name:    "Host",
			Email:   "host@user.com",
			Address: "127.0.0.1:8082",
			Online:  true,
		}

		// Marshal the host peer into JSON
		hostJSON, err := json.Marshal(host)
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP request with the host JSON as the request body
		req, err := http.NewRequest("POST", "/rooms", bytes.NewBuffer(hostJSON))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Create a new HTTP recorder
		w := httptest.NewRecorder()

		// Serve the HTTP request to the router
		router.ServeHTTP(w, req)

		// Check the HTTP status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Assert that the response body contains the expected JSON with the room_id
		var responseBody map[string]string
		err = json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Contains(t, responseBody, "room_id")
	})

	// Test case 2: Error handling when binding JSON fails
	t.Run("BindingJSONError", func(t *testing.T) {
		// Create a new HTTP request with an invalid JSON payload
		req, err := http.NewRequest("POST", "/rooms", bytes.NewBuffer([]byte("invalid JSON")))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Create a new HTTP recorder
		w := httptest.NewRecorder()

		// Serve the HTTP request to the router
		router.ServeHTTP(w, req)

		// Check the HTTP status code
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// Assert that the response body contains the error message
		assert.Contains(t, w.Body.String(), "invalid character 'i' looking for beginning of value")
	})
	// Test case 3: Error handling when room creation fails
	t.Run("ErrorHandlingRoomCreationFailure", func(t *testing.T) {
		// Simulate an error in CreateRoom function by passing a host with empty required fields
		host := &network.Peer{}
		_, err := transport.CreateRoom(host)

		// Check if the error is not nil
		assert.Error(t, err)

		hostJSON, err := json.Marshal(host)
		if err != nil {
			t.Fatal(err)
		}

		// Create a new HTTP request with the host JSON as the request body
		req, err := http.NewRequest("POST", "/rooms", bytes.NewBuffer(hostJSON))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Create a new HTTP recorder
		w := httptest.NewRecorder()

		// Serve the HTTP request to the router
		router.ServeHTTP(w, req)

		// Check the HTTP status code
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		// Assert that the response body contains the error message
		assert.Contains(t, w.Body.String(), "host peer must have ID, Name, Email and Address fields")
	})

}

func TestSendChatMessage(t *testing.T) {
	// Create a new instance of TCPTransport and ChatService
	transport := network.NewTCPTransport()
	chatService := chat.NewChatService(transport)

	// Create a new chat controller
	chatController := NewChatController(transport, chatService)

	// Create a new room
	host := &network.Peer{
		ID:      "host123",
		Name:    "Host",
		Address: "host@example.com",
		Email:   "host@example.com",
	}
	roomID, err := transport.CreateRoom(host)
	if err != nil {
		t.Fatalf("failed to create room: %v", err)
	}

	// Create a request body with sender ID and message
	requestBody := map[string]string{
		"sender_id": "sender123",
		"message":   "Test message",
	}
	requestBodyBytes, _ := json.Marshal(requestBody)

	// Create a new request
	req, err := http.NewRequest("POST", fmt.Sprintf("/rooms/%s/chat", roomID), bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a new recorder
	recorder := httptest.NewRecorder()

	// Set up the Gin context
	context, _ := gin.CreateTestContext(recorder)
	context.Params = gin.Params{
		{Key: "roomID", Value: roomID},
	}

	// Bind request body to context
	context.Request = req

	// Call the controller method
	chatController.SendChatMessage(context)

	// Verify the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body
	var responseBody map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	// Check if the response contains the expected message
	assert.Contains(t, responseBody, "message")
	assert.Equal(t, "Message sent successfully", responseBody["message"])

	// Check if the response contains the chat history
	assert.Contains(t, responseBody, "chat_history")
	chatHistory, ok := responseBody["chat_history"].([]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, chatHistory)
}
