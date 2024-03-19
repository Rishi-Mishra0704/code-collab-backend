package filefolder

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// ReadFileContent handles HTTP requests to read the content of a file.
// It accepts a POST request with a JSON body containing the path of the file whose content is to be read.
// The request body should have the following structure:
//
//	{
//	    "path": "string"  // Path to the file whose content is to be read
//	}
//
// Upon receiving the request, it reads the content of the specified file and returns a JSON response with HTTP status 200 (OK)
// containing the content of the file as a string.
// In case of any errors during the process, it returns a JSON response with HTTP status 500 (Internal Server Error)
// along with an error message providing details of the encountered error.
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
