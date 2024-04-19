// controllers/chat_controller.go

package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Rishi-Mishra0704/code-collab-backend/chat"
	"github.com/Rishi-Mishra0704/code-collab-backend/network"
)

// ChatController represents the controller for chat-related endpoints.
type ChatController struct {
	TCPTransport *network.TCPTransport // Reference to the TCPTransport instance
	ChatService  *chat.ChatService     // Reference to the ChatService instance
}

// NewChatController creates a new instance of ChatController.
func NewChatController(transport *network.TCPTransport, chatService *chat.ChatService) *ChatController {
	return &ChatController{
		TCPTransport: transport,
		ChatService:  chatService,
	}
}

// CreateRoom handles the creation of a new chat room.
func (cc *ChatController) CreateRoom(c *gin.Context) {
	// Parse request body to get host peer details
	var host network.Peer
	if err := c.BindJSON(&host); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create room and get room ID
	roomID, err := cc.TCPTransport.CreateRoom(&host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"room_id": roomID})
}

// JoinRoom handles a peer joining an existing chat room.
func (cc *ChatController) JoinRoom(c *gin.Context) {
	// Parse request body to get peer details including room_id
	var request struct {
		RoomID string `json:"room_id"`
		Peer   network.Peer
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Join room with peer
	err := cc.TCPTransport.JoinRoom(request.RoomID, &request.Peer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Peer %s joined room %s", request.Peer.ID, request.RoomID)})
}

// LeaveRoom handles a peer leaving a chat room.
// Update the LeaveRoom function to properly extract the peer ID from the request parameters.
func (cc *ChatController) LeaveRoom(c *gin.Context) {
	roomID := c.Param("roomID")
	peerID := c.Param("peerID") // Ensure peerID is correctly extracted

	// Leave room
	err := cc.TCPTransport.LeaveRoom(roomID, peerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Peer %s left room %s", peerID, roomID)})
}

func (cc *ChatController) GetRooms(c *gin.Context) {
	// Get all rooms from the TCPTransport
	rooms := cc.TCPTransport.GetAllRooms()

	// Return the rooms as JSON response
	c.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

// SendChatMessage handles sending a chat message to a room.
func (cc *ChatController) SendChatMessage(c *gin.Context) {
	roomID := c.Param("roomID")

	// Parse request body to get sender and message details
	var message struct {
		SenderID string `json:"sender_id"`
		Message  string `json:"message"`
	}
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the sender peer from the request parameters
	senderID := message.SenderID

	// Create a new peer with the sender ID
	sender := &network.Peer{
		ID: senderID,
	}

	// Send the message to the room using the ChatService
	err := cc.ChatService.Send(roomID, sender, message.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Receive the updated chat history after sending the message
	chatHistory, err := cc.ChatService.Receive(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully", "chat_history": chatHistory})
}
