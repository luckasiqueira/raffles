package main

import (
	"embed"
	"raffles/internal/server"
)

//go:embed internal/frontend/assets
var fs embed.FS

// Entrypoint of entire application, which calls the start of webserver
func main() {
	server.StartServer(fs)
}
