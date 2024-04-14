package server

import (
	"log"
	"net/http"

	"github.com/Rishi-Mishra0704/code-collab-backend/collab"
	"github.com/Rishi-Mishra0704/code-collab-backend/compiler"
	"github.com/Rishi-Mishra0704/code-collab-backend/controllers"
	"github.com/gorilla/handlers"
)

func StartHTTPServer() {
	// Initialize WebSocket router
	wsRouter := http.NewServeMux()
	wsRouter.HandleFunc("/collab", collab.HandleCollaborations)
	wsRouter.HandleFunc("/execute", controllers.ExecuteCommand)
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
