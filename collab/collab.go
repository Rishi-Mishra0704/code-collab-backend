package collab

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	models "github.com/Rishi-Mishra0704/code-collab-backend/models"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan models.File)       // broadcast channel

// Configure the WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleCollaborations(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Register new client
	clients[ws] = true

	for {
		var file models.File
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&file)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- file
	}
}

func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
