package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Display prompt to the client
	fmt.Fprint(conn, "127.0.0.1:1234> ")

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
			fmt.Fprint(conn, "Invalid command\n")
			continue
		}

		command := strings.ToUpper(parts[0])

		switch command {
		case "SET":
			// Allow for multiple values (at least a key and one value)
			if len(parts) < 3 {
				fmt.Fprint(conn, "Usage: SET key value1 value2 ...\n")
				continue
			}
			key := parts[1]
			values := parts[2:] // Capture all values after the key
			err := s.store.Set(key, values...)
			if err != nil {
				fmt.Fprint(conn, err.Error()+"\n")
			} else {
				fmt.Fprint(conn, "OK\n")
			}

		case "GET":
			if len(parts) != 2 {
				fmt.Fprint(conn, "Usage: GET key\n")
				continue
			}
			key := parts[1]
			values, err := s.store.Get(key)
			if err != nil {
				fmt.Fprint(conn, err.Error()+"\n")
			} else {
				fmt.Fprint(conn, "Values: "+strings.Join(values, ", ")+"\n")
			}

		case "DELETE":
			if len(parts) != 2 {
				fmt.Fprint(conn, "Usage: DELETE key\n")
				continue
			}
			key := parts[1]
			err := s.store.Delete(key)
			if err != nil {
				fmt.Fprint(conn, err.Error()+"\n")
			} else {
				fmt.Fprint(conn, "Deleted\n")
			}

		case "EXISTS":
			if len(parts) != 2 {
				fmt.Fprint(conn, "Usage: EXISTS key\n")
				continue
			}
			key := parts[1]
			exists := s.store.Exists(key)
			if exists {
				fmt.Fprint(conn, "1\n")
			} else {
				fmt.Fprint(conn, "0\n")
			}

		case "EXPIRE":
			if len(parts) != 3 {
				fmt.Fprint(conn, "Usage: EXPIRE key ttl_in_seconds\n")
				continue
			}
			key := parts[1]
			ttl, err := time.ParseDuration(parts[2] + "s")
			if err != nil {
				fmt.Fprint(conn, "Invalid TTL value\n")
				continue
			}
			err = s.store.Expire(key, ttl)
			if err != nil {
				fmt.Fprint(conn, err.Error()+"\n")
			} else {
				fmt.Fprint(conn, "OK\n")
			}

		case "TTL":
			if len(parts) != 2 {
				fmt.Fprint(conn, "Usage: TTL key\n")
				continue
			}
			key := parts[1]
			ttl, err := s.store.TTL(key)
			if err != nil {
				fmt.Fprint(conn, err.Error()+"\n")
			} else {
				fmt.Fprintf(conn, "TTL: %d seconds\n", int(ttl.Seconds()))
			}

		case "RAND":
			if len(parts) != 3 {
				fmt.Fprint(conn, "Usage: RAND key offset\n")
				continue
			}
			key := parts[1]
			offset, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Fprint(conn, "Invalid offset value\n")
				continue
			}
			value, err := s.store.RandomValues(key, offset)
			if err != nil {
				fmt.Fprint(conn, err.Error()+"\n")
			} else {
				fmt.Fprint(conn, "Values: "+strings.Join(value, ". ")+"\n")
			}

		case "EXIT":
			fmt.Fprint(conn, "Bye!\n")
			conn.Close()
			return

		case "HELP":
			fmt.Fprint(conn, "Available commands:\n")
			fmt.Fprint(conn, "SET key value\n")
			fmt.Fprint(conn, "GET key\n")
			fmt.Fprint(conn, "DELETE key\n")
			fmt.Fprint(conn, "EXISTS key\n")
			fmt.Fprint(conn, "EXPIRE key ttl_in_seconds\n")
			fmt.Fprint(conn, "TTL key\n")
			fmt.Fprint(conn, "EXIT\n")
			fmt.Fprint(conn, "HELP\n")

		default:
			fmt.Fprint(conn, "Unknown command\n")
		}

		// Print prompt after handling the command
		fmt.Fprint(conn, "127.0.0.1:1234> ")
	}
}
