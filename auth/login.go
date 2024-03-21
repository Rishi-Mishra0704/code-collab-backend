package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginRequest struct {
		Name     string `json:"name"`     // Username provided in the login request
		Password string `json:"password"` // Password provided in the login request
	}
	// Parse the JSON request body into the loginRequest struct.
	if err := c.BindJSON(&loginRequest); err != nil {
		// Respond with an error if the request payload is invalid.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Check if the user exists in the users map and if the provided password matches the stored password.
	user, exists := peers[loginRequest.Name]
	if !exists {
		// Respond with an error if the username or password is invalid.
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Respond with a success message if the login process is successful.
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "user": user})
}
