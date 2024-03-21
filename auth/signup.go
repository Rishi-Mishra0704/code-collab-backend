package auth

import (
	"net/http"

	"github.com/Rishi-Mishra0704/code-collab-backend/network"
	"github.com/gin-gonic/gin"
)

var peers = map[string]network.Peer{}

// SignupHandler handles the signup process for a new peer.
func SignupHandler(c *gin.Context) {
	var newPeer network.Peer
	// Parse the JSON request body into the newPeer struct.
	if err := c.BindJSON(&newPeer); err != nil {
		// Respond with an error if the request payload is invalid.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Check if the peer ID already exists in the peers map. If it does, return an error.
	if _, exists := peers[newPeer.ID]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Peer ID already exists"})
		return
	}
	// Check if the email already exists among existing peers. If it does, return an error.
	for _, existingPeer := range peers {
		if existingPeer.Email == newPeer.Email {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}
	}

	// Add the new peer to the peers map.
	peers[newPeer.ID] = newPeer
	// Respond with a success message if the signup process is successful.
	c.JSON(http.StatusOK, gin.H{"message": "Peer signed up successfully"})
}
