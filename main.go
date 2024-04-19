package main

import (
	"log"
	"net/http"

	"github.com/Rishi-Mishra0704/code-collab-backend/chat"
	"github.com/Rishi-Mishra0704/code-collab-backend/collab"
	"github.com/Rishi-Mishra0704/code-collab-backend/compiler"
	"github.com/Rishi-Mishra0704/code-collab-backend/controllers"
	filefolder "github.com/Rishi-Mishra0704/code-collab-backend/file-folder"
	"github.com/Rishi-Mishra0704/code-collab-backend/network"
	"github.com/gin-contrib/cors"
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
	apiRouter.Use(cors.Default())
	// Define API endpoints using the ChatController methods

	// Room operations
	apiRouter.GET("/rooms", chatController.GetRooms)
	apiRouter.POST("/create-room", chatController.CreateRoom)
	apiRouter.POST("/join-room/", chatController.JoinRoom)
	apiRouter.POST("/leave-room/:roomID/:peerID", chatController.LeaveRoom)
	apiRouter.POST("/rooms/:roomID/send-message", chatController.SendChatMessage)
	apiRouter.GET("/rooms/:roomID/chats", chatController.GetChatHistory)
	// File and folder operations
	apiRouter.POST("create", filefolder.CreateFileOrFolder)
	apiRouter.POST("list", filefolder.ListFilesOrFolder)
	apiRouter.POST("read", filefolder.ReadFileContent)
	// Start Gin server for REST API
	go func() {
		if err := apiRouter.Run(":8080"); err != nil {
			log.Fatalf("Failed to start API server: %v", err)
		}
	}()

	// Initialize WebSocket router
	wsRouter := http.NewServeMux()
	// handle Collaborations
	wsRouter.HandleFunc("/collab", collab.HandleCollaborations)
	// Execute terminal commands
	wsRouter.HandleFunc("/execute", controllers.ExecuteCommand)
	// Execute code
	wsRouter.HandleFunc("/compile", compiler.ExecuteCodeHandler)
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
