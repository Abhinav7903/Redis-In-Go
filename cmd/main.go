package main

import (
	"go-idis/internal/idis"
	"go-idis/server"
	"log"
	"time"
)

// main is the entry point of the application. It initializes and runs an in-memory key-value store server
// with both HTTP and Telnet interfaces. The server includes the following features:
//
// - Creates an in-memory repository for storing key-value pairs
// - Starts HTTP server on 0.0.0.0:1234
// - Starts Telnet server on 0.0.0.0:5678
// - Implements an automatic cleanup mechanism that deletes all keys after 2 minutes of server start
// - Sets up periodic data persistence by dumping the store contents to 'dump.json' every 2 hours
//
// The server runs until an error occurs or the process is terminated.
// If the server encounters a fatal error, it will log the error and terminate the program.

func main() {
	// Initialize the	 in-memory repository
	store := idis.NewInMemoryRepository()

	// Create a new server instance
	httpAddr := "0.0.0.0:1234"   // HTTP server address
	telnetAddr := "0.0.0.0:5678" // Telnet server address
	srv := server.NewServer(httpAddr, telnetAddr, store)

	// Goroutine to delete all keys after 5 minutes of server start (for because of deploying on Internet)
	go func() {
		// Wait for 5 minutes
		time.Sleep(2 * time.Minute)

		// Delete all keys
		err := store.DeleteAll()
		if err != nil {
			log.Printf("Error deleting all keys: %v", err)
		} else {
			log.Println("All keys deleted successfully after 2 minutes.")
		}
	}()

	// Set up periodic data dump to a file
	filepath := "dump.json"
	store.StartAutoDump(filepath, 2*time.Hour)

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
