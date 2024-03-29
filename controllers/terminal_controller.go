package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/Rishi-Mishra0704/code-collab-backend/console"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ExecuteCommand(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		_, cmdBytes, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading command: %v", err)
			return
		}

		command := string(cmdBytes)

		if strings.TrimSpace(command) == "" {
			continue
		}

		// Execute command
		output, err := console.CallTerminal(command)
		if err != nil {
			log.Printf("Error executing command: %v", err)
			continue
		}

		// Send output back to client
		if err := conn.WriteMessage(websocket.TextMessage, []byte(output)); err != nil {
			log.Printf("Error sending output: %v", err)
			return
		}
	}
}
