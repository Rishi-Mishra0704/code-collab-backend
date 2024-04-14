package main

import (
	"github.com/Rishi-Mishra0704/code-collab-backend/server"
)

func main() {
	server.StartGinServer()
	server.StartHTTPServer()
}
