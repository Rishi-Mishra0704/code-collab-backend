package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Rishi-Mishra0704/code-collab-backend/chat"
	"github.com/Rishi-Mishra0704/code-collab-backend/controllers"
	"github.com/Rishi-Mishra0704/code-collab-backend/network"
)

func main() {
	// Initialize TCP transport
	transport := network.NewTCPTransport()
	chatService := chat.NewChatService(transport)

	// Initialize ChatController with ChatServiceport
	chatController := controllers.NewChatController(transport, chatService)

	// Initialize Gin router
	router := gin.Default()

	// Define API endpoints using the ChatController methods
	router.GET("/rooms", chatController.GetRooms)
	router.POST("/create-room", chatController.CreateRoom)
	router.POST("/join-room/:roomID", chatController.JoinRoom)
	router.POST("/leave-room/:roomID/:peerID", chatController.LeaveRoom)
	router.POST("/rooms/:roomID/chat", chatController.SendChatMessage)
	// Start Gin server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
