package server

import (
	"fmt"
	"net"
)

func (s *Server) handleSet(conn net.Conn, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: SET key value1 value2.... valueN")
	}
	key, values := args[0], args[1:]
	if err := s.store.Set(key, values...); err != nil {
		return err
	}
	fmt.Fprint(conn, "OK\n")
	return nil
}

func (s *Server) handleSetUnique(conn net.Conn, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: SETUQ key value1 value2 ... valueN")
	}
	key, values := args[0], args[1:]
	if err := s.store.SetUnique(key, values...); err != nil {
		return err
	}
	fmt.Fprint(conn, "OK\n")
	return nil
}
