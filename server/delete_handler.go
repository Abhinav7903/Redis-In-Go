package server

import (
	"fmt"
	"net"
)

func (s *Server) handleDelete(conn net.Conn, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: DELETE key")
	}
	key := args[0]
	if err := s.store.Delete(key); err != nil {
		return err
	}
	fmt.Fprint(conn, "Deleted\n")
	return nil
}
