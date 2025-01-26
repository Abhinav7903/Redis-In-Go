package main

import (
	"go-idis/internal/idis"
	"go-idis/server"
	"log"
	"time"
)

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
		time.Sleep(5 * time.Minute)

		// Delete all keys
		err := store.DeleteAll()
		if err != nil {
			log.Printf("Error deleting all keys: %v", err)
		} else {
			log.Println("All keys deleted successfully after 5 minutes.")
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
