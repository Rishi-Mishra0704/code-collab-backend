package controllers

import (
	"net/http"

	"github.com/Rishi-Mishra0704/code-collab-backend/console"
	"github.com/gin-gonic/gin"
)

type CommandRequest struct {
	Command string `json:"command"`
}

func ExecuteCommand(c *gin.Context) {
	var req CommandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if req.Command == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Command parameter is required"})
		return
	}

	output, err := console.CallTerminal(req.Command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, output)
}
