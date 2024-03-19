package filefolder

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// ReadFileContent handles HTTP requests to read the content of a file.
func ReadFileContent(c *gin.Context) {
	// Define a struct to hold request parameters.
	var req struct {
		Path string `json:"path"` // Path to the file whose content is to be read
	}

	// Bind request body to the req struct.
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Read the content of the file.
	fileContent, err := os.ReadFile(req.Path)
	if err != nil {
		// If an error occurred while reading the file, return an internal server error response.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file content", "detail": err.Error()})
		return
	}

	// Return the content of the file.
	c.JSON(http.StatusOK, gin.H{"content": string(fileContent)})
}
