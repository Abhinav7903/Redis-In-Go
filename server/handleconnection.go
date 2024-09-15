package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	prompt := "127.0.0.1:1234> "

	for {
		// Display prompt to the client
		fmt.Fprint(conn, prompt)

		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		// Process command
		if err := s.processCommand(conn, strings.TrimSpace(message)); err != nil {
			fmt.Fprint(conn, err.Error()+"\n")
		}
	}
}

func (s *Server) processCommand(conn net.Conn, message string) error {
	parts := strings.Split(message, " ")

	if len(parts) == 0 {
		return fmt.Errorf("invalid command")
	}

	command := strings.ToUpper(parts[0])
	args := parts[1:]

	switch command {
	case "SET":
		return s.handleSet(conn, args)
	case "GET":
		return s.handleGet(conn, args)
	case "DELETE":
		return s.handleDelete(conn, args)
	case "EXISTS":
		return s.handleExists(conn, args)
	case "EXPIRE":
		return s.handleExpire(conn, args)
	case "TTL":
		return s.handleTTL(conn, args)
	case "RAND":
		return s.handleRand(conn, args)
	case "SETUQ":
		return s.handleSetUnique(conn, args)
	case "REMOVE":
		return s.handleRemove(conn, args)
	case "GETUQ":
		return s.handleGetUnique(conn, args)
	case "GETKEY":
		return s.handleGetKey(conn, args)
	case "EXIT":
		fmt.Fprint(conn, "Goodbye!\n")
		conn.Close()
		return nil
	case "HELP":
		return s.handleHelp(conn)
	default:
		return fmt.Errorf("unknown command")
	}
}
