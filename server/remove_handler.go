package server

import (
	"fmt"
	"net"
)

func (s *Server) handleRemove(conn net.Conn, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: REMOVE key value")
	}
	key, value := args[0], args[1]
	if err := s.store.RemoveValue(key, value); err != nil {
		return err
	}
	fmt.Fprint(conn, "Removed\n")
	return nil
}
