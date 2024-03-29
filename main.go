package main

import (
	"log"
	"net/http"

	"github.com/Rishi-Mishra0704/code-collab-backend/chat"
	"github.com/Rishi-Mishra0704/code-collab-backend/controllers"
	"github.com/Rishi-Mishra0704/code-collab-backend/network"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

func main() {
	// Initialize TCP transport
	transport := network.NewTCPTransport()
	chatService := chat.NewChatService(transport)

	// Initialize ChatController with ChatService
	chatController := controllers.NewChatController(transport, chatService)

	// Initialize Gin router for REST API
	apiRouter := gin.Default()

	// Define API endpoints using the ChatController methods
	apiRouter.GET("/rooms", chatController.GetRooms)
	apiRouter.POST("/create-room", chatController.CreateRoom)
	apiRouter.POST("/join-room/:roomID", chatController.JoinRoom)
	apiRouter.POST("/leave-room/:roomID/:peerID", chatController.LeaveRoom)
	apiRouter.POST("/rooms/:roomID/chat", chatController.SendChatMessage)

	// Start Gin server for REST API
	go func() {
		if err := apiRouter.Run(":8080"); err != nil {
			log.Fatalf("Failed to start API server: %v", err)
		}
	}()

	// Initialize WebSocket router
	wsRouter := http.NewServeMux()
	wsRouter.HandleFunc("/execute", controllers.ExecuteCommand)

	// Apply CORS middleware to the WebSocket server
	wsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(wsRouter)

	// Start WebSocket server
	if err := http.ListenAndServe(":8000", wsHandler); err != nil {
		log.Fatalf("Failed to start WebSocket server: %v", err)
	}
}
