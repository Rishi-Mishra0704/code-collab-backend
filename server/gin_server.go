package server

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Rishi-Mishra0704/code-collab-backend/chat"
	"github.com/Rishi-Mishra0704/code-collab-backend/controllers"
	filefolder "github.com/Rishi-Mishra0704/code-collab-backend/file-folder"
	"github.com/Rishi-Mishra0704/code-collab-backend/network"
)

func StartGinServer() {
	// Initialize TCP transport
	transport := network.NewTCPTransport()
	chatService := chat.NewChatService(transport)

	// Initialize ChatController with ChatService
	chatController := controllers.NewChatController(transport, chatService)

	// Initialize Gin router for REST API
	apiRouter := gin.Default()
	apiRouter.Use(cors.Default())
	// Define API endpoints using the ChatController methods
	apiRouter.GET("/rooms", chatController.GetRooms)
	apiRouter.POST("/create-room", chatController.CreateRoom)
	apiRouter.POST("/join-room/:roomID", chatController.JoinRoom)
	apiRouter.POST("/leave-room/:roomID/:peerID", chatController.LeaveRoom)
	apiRouter.POST("/rooms/:roomID/chat", chatController.SendChatMessage)
	apiRouter.POST("create", filefolder.CreateFileOrFolder)
	apiRouter.POST("list", filefolder.ListFilesOrFolder)
	apiRouter.POST("read", filefolder.ReadFileContent)
	// Start Gin server for REST API
	go func() {
		if err := apiRouter.Run(":8080"); err != nil {
			log.Fatalf("Failed to start API server: %v", err)
		}
	}()
}
