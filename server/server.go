package server

import (
	"fmt"
	"go-idis/internal/idis"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	httpAddr   string
	telnetAddr string
	store      idis.Repository
	router     *mux.Router
}

// NewServer initializes the Server with HTTP and Telnet addresses
func NewServer(httpAddr, telnetAddr string, store idis.Repository) *Server {
	return &Server{
		httpAddr:   httpAddr,
		telnetAddr: telnetAddr,
		store:      store,
		router:     mux.NewRouter(),
	}
}

// Run starts the HTTP and Telnet servers concurrently
func (s *Server) Run() error {
	// Start the HTTP server in a separate goroutine
	go func() {
		s.RegisterAPIs()
		fmt.Printf("HTTP server running on %s\n", s.httpAddr)
		if err := http.ListenAndServe(s.httpAddr, s.router); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// Start the Telnet server
	listener, err := net.Listen("tcp", s.telnetAddr)
	if err != nil {
		return fmt.Errorf("telnet server failed to start: %w", err)
	}
	defer listener.Close()
	fmt.Printf("Telnet server running on %s\n", s.telnetAddr)
	fmt.Println("Type 'exit' to shut down the Telnet server.")

	// Accept Telnet connections in a loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		fmt.Printf("Telnet client connected from %s\n", conn.RemoteAddr().String())
		go s.handleConnection(conn)
	}
}
