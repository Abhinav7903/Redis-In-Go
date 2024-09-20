package server

import (
	"fmt"
	"net"
)

func (s *Server) handleLoadDump(conn net.Conn, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: LOADDUMP filepath")
	}
	filepath := args[0]
	err := s.store.LoadFromDump(filepath)
	if err != nil {
		return err
	}
	fmt.Fprintf(conn, "Data successfully loaded from file: %s\n", filepath)
	return nil
}
