// controllers/chat_controller.go

package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Rishi-Mishra0704/code-collab-backend/network"
)

// ChatController represents the controller for chat-related endpoints.
type ChatController struct {
	TCPTransport *network.TCPTransport // Reference to the TCPTransport instance
}

// NewChatController creates a new instance of ChatController.
func NewChatController(transport *network.TCPTransport) *ChatController {
	return &ChatController{
		TCPTransport: transport,
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
	roomID := c.Param("roomID")

	// Parse request body to get peer details
	var peer network.Peer
	if err := c.BindJSON(&peer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Join room with peer
	err := cc.TCPTransport.JoinRoom(roomID, &peer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Peer %s joined room %s", peer.ID, roomID)})
}

// LeaveRoom handles a peer leaving a chat room.
func (cc *ChatController) LeaveRoom(c *gin.Context) {
	roomID := c.Param("roomID")
	peerID := c.Param("peerID")

	// Leave room
	err := cc.TCPTransport.LeaveRoom(roomID, peerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Peer %s left room %s", peerID, roomID)})
}
