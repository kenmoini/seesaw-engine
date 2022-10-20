package main

import (
	httpserver "github.com/kenmoini/seesaw-server/internal/server"
)

// main is the entrypoint for the application, should be as small as possible
func main() {
	// Run the application
	httpserver.StartServer()
}
