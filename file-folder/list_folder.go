package filefolder

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// ListFilesOrFolder handles HTTP requests to list files or folders in a directory.
// It accepts a POST request with a JSON body containing the path of the directory whose files or folders are to be listed.
// The request body should have the following structure:
//
//	{
//	    "path": "string"  // Path to the directory whose files or folders are to be listed
//	}
//
// Upon receiving the request, it reads the contents of the specified directory and returns a JSON response with HTTP status 200 (OK)
// containing an array of objects, each representing a file or folder, with properties 'name' and 'type'.
// 'name' represents the name of the file or folder, and 'type' indicates whether it is a file or a folder.
// In case of any errors during the process, it returns a JSON response with HTTP status 500 (Internal Server Error)
// and an error message.
func ListFilesOrFolder(c *gin.Context) {
	// Define a struct to hold request parameters.
	var req struct {
		Path string `json:"path"` // Path to the directory whose files or folders are to be listed
	}

	// Bind request body to the req struct.
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the received path for debugging.
	fmt.Println("Received path:", req.Path)

	// Read directory contents.
	files, err := os.ReadDir(req.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a list to hold file and folder objects.
	var fileList []map[string]interface{}
	for _, file := range files {
		// Initialize a map to represent each file or folder.
		item := make(map[string]interface{})
		item["name"] = file.Name()
		// Determine the type of the item (file or folder).
		if file.IsDir() {
			item["type"] = "folder"
		} else {
			item["type"] = "file"
		}
		// Append the item to the list.
		fileList = append(fileList, item)
	}

	// Return the list of file and folder objects in the directory.
	c.JSON(http.StatusOK, gin.H{"files": fileList})
}
