package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents the structure of a user in the system.
type User struct {
	Username string `json:"username"` // Unique username of the user
	Email    string `json:"email"`    // Email address of the user
	Password string `json:"password"` // Password of the user
}

// users is a map that stores user information with the username as the key.
var users = map[string]User{}

// SignupHandler handles the signup process for a new user.
func SignupHandler(c *gin.Context) {
	var newUser User
	// Parse the JSON request body into the newUser struct.
	if err := c.BindJSON(&newUser); err != nil {
		// Respond with an error if the request payload is invalid.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Check if the username already exists in the users map. If it does, return an error.
	if _, exists := users[newUser.Username]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	// Check if the email already exists among existing users. If it does, return an error.
	for _, existingUser := range users {
		if existingUser.Email == newUser.Email {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}
	}

	// Add the new user to the users map.
	users[newUser.Username] = newUser
	// Respond with a success message if the signup process is successful.
	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully"})
}

// LoginHandler handles the login process for existing users.
func LoginHandler(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"` // Username provided in the login request
		Password string `json:"password"` // Password provided in the login request
	}
	// Parse the JSON request body into the loginRequest struct.
	if err := c.BindJSON(&loginRequest); err != nil {
		// Respond with an error if the request payload is invalid.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Check if the user exists in the users map and if the provided password matches the stored password.
	user, exists := users[loginRequest.Username]
	if !exists || user.Password != loginRequest.Password {
		// Respond with an error if the username or password is invalid.
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Respond with a success message if the login process is successful.
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}
