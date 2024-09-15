package server

import (
	"fmt"
	"go-idis/internal/idis"
	"log"
	"net"
)

type Server struct {
	addr  string
	store idis.Repository
}

func NewServer(addr string, store idis.Repository) *Server {
	return &Server{
		addr:  addr,
		store: store,
	}
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Printf("Server running on %s\n", s.addr)
	fmt.Println("Type 'exit' to shut down the server.")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}

		fmt.Printf("Client connected from %s\n", conn.RemoteAddr().String())
		go s.handleConnection(conn)
	}
}
