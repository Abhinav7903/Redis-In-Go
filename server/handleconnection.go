package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		// Trim newline and split the command
		message = strings.TrimSpace(message)
		parts := strings.Split(message, " ")

		if len(parts) == 0 {
			conn.Write([]byte("Invalid command\n"))
			continue
		}

		command := strings.ToUpper(parts[0])

		switch command {
		case "SET":
			if len(parts) != 3 {
				conn.Write([]byte("Usage: SET key value\n"))
				continue
			}
			key, value := parts[1], parts[2]
			err := s.store.Set(key, value)
			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
			} else {
				conn.Write([]byte("OK\n"))
			}
		case "GET":
			if len(parts) != 2 {
				conn.Write([]byte("Usage: GET key\n"))
				continue
			}
			key := parts[1]
			value, err := s.store.Get(key)
			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
			} else {
				conn.Write([]byte("Value: " + value + "\n"))
			}
		case "DELETE":
			if len(parts) != 2 {
				conn.Write([]byte("Usage: DELETE key\n"))
				continue
			}
			key := parts[1]
			err := s.store.Delete(key)
			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
			} else {
				conn.Write([]byte("Deleted\n"))
			}
		case "EXISTS":
			if len(parts) != 2 {
				conn.Write([]byte("Usage: EXISTS key\n"))
				continue
			}
			key := parts[1]
			exists := s.store.Exists(key)
			if exists {
				conn.Write([]byte("1\n"))
			} else {
				conn.Write([]byte("0\n"))
			}
		case "EXPIRE":
			if len(parts) != 3 {
				conn.Write([]byte("Usage: EXPIRE key ttl_in_seconds\n"))
				continue
			}
			key := parts[1]
			ttl, err := time.ParseDuration(parts[2] + "s")
			if err != nil {
				conn.Write([]byte("Invalid TTL value\n"))
				continue
			}
			err = s.store.Expire(key, ttl)
			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
			} else {
				conn.Write([]byte("OK\n"))
			}
		case "TTL":
			if len(parts) != 2 {
				conn.Write([]byte("Usage: TTL key\n"))
				continue
			}
			key := parts[1]
			ttl, err := s.store.TTL(key)
			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
			} else {
				conn.Write([]byte(fmt.Sprintf("TTL: %d seconds\n", int(ttl.Seconds()))))
			}

		case "Exit":
			conn.Write([]byte("Bye!\n"))
			return
		case "HELP":
			conn.Write([]byte("Available commands:\n"))
			conn.Write([]byte("SET key value\n"))
			conn.Write([]byte("GET key\n"))
			conn.Write([]byte("DELETE key\n"))
			conn.Write([]byte("EXISTS key\n"))
			conn.Write([]byte("EXPIRE key ttl_in_seconds\n"))
			conn.Write([]byte("TTL key\n"))
			conn.Write([]byte("EXIT\n"))
			conn.Write([]byte("HELP\n"))
		default:
			conn.Write([]byte("Unknown command\n"))
		}
	}
}
