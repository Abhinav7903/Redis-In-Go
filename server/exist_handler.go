package server

import (
	"fmt"
	"net"
)

func (s *Server) handleExists(conn net.Conn, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: EXISTS key")
	}
	key := args[0]
	exists := s.store.Exists(key)
	if exists {
		fmt.Fprint(conn, "1\n")
	} else {
		fmt.Fprint(conn, "0\n")
	}
	return nil
}
