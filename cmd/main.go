package main

import (
	"go-idis/internal/idis"
	"go-idis/server"
	"log"
)

func main() {
	// Initialize the in-memory repository
	store := idis.NewInMemoryRepository()

	// Create a new TCP server
	srv := server.NewServer("0.0.0.0:1234", store) //you can change the address and port here to your desired address and port

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
