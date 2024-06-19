package main

import "raffles/internal/server"

// Entrypoint of entire application, which calls the start of webserver
func main() {
	server.StartServer()
}
