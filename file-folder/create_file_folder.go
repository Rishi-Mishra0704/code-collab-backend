package filefolder

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// CreateFileOrFolder handles HTTP requests to create a file or folder.
// It accepts a POST request with a JSON body containing the path and type of file/folder to be created.
// The request body should have the following structure:
//
//	{
//	    "path": "string",     // Path to the file or folder to be created
//	    "isFolder": "bool"    // Indicates whether the request is to create a folder
//	}
//
// If 'isFolder' is true, it will create a folder at the specified path.
// If 'isFolder' is false, it will create a file at the specified path.
// Upon successful creation, it returns a JSON response with HTTP status 201 (Created) and a success message.
// In case of any errors during the creation process, it returns a JSON response with HTTP status 500 (Internal Server Error)
// and an error message.
func CreateFileOrFolder(c *gin.Context) {
	// Define a struct to hold request parameters.
	var req struct {
		Path     string `json:"path"`     // Path to the file or folder to be created
		IsFolder bool   `json:"isFolder"` // Indicates whether the request is to create a folder
	}

	// Bind request body to the req struct.
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Declare an error variable.
	var err error

	// If IsFolder is true, create a folder with the specified path.
	if req.IsFolder {
		err = os.MkdirAll(req.Path, 0755)
	} else {
		// If IsFolder is false, create a file with the specified path.
		_, err = os.Create(req.Path)
	}

	// If an error occurred during file or folder creation, return an internal server error response.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If the file or folder was created successfully, return a success response.
	c.JSON(http.StatusCreated, gin.H{"message": "File or folder created successfully"})
}
