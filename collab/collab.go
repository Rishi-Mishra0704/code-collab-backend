package collab

import (
	"net/http"

	"github.com/gorilla/websocket"

	models "github.com/Rishi-Mishra0704/code-collab-backend/models"
)

/*
1. Initialization:
   - A map called 'clients' is created to store WebSocket connections.
   - A channel called 'broadcast' is created to broadcast messages to all connected clients.

2. Struct Definition:
   - A struct named 'File' is defined to represent file data, including content and file extension.

3. Main Function:
   - Configures the WebSocket route '/ws'.
   - Starts listening for incoming connections.
   - Sets up Cross-Origin Resource Sharing (CORS) to allow requests from any origin.

4. handleConnections Function:
   - Accepts incoming HTTP requests and upgrades them to WebSocket connections.
   - Registers clients in the 'clients' map.
   - Reads incoming JSON messages representing files from clients.
   - Broadcasts the received file messages to all connected clients.

5. handleMessages Function:
   - Continuously listens for messages on the 'broadcast' channel.
   - Sends each received message to all connected clients.

6. WebSocket Upgrader Configuration:
   - Configures the WebSocket upgrader with a custom 'CheckOrigin' function allowing connections from any origin.

7. CORS Configuration:
   - Sets up Cross-Origin Resource Sharing (CORS) to allow requests from any origin, with specific allowed methods and headers.
*/

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan models.File)       // broadcast channel

// Configure the WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
