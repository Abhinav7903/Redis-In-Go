package server

import (
	"fmt"
	"net"
	"time"
)

func (s *Server) handleExpire(conn net.Conn, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: EXPIRE key ttl_in_seconds")
	}
	key, ttlStr := args[0], args[1]
	ttl, err := time.ParseDuration(ttlStr + "s")
	if err != nil {
		return fmt.Errorf("invalid TTL value")
	}
	if err := s.store.Expire(key, ttl); err != nil {
		return err
	}
	fmt.Fprint(conn, "OK\n")
	return nil
}

func (s *Server) handleTTL(conn net.Conn, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: TTL key")
	}
	key := args[0]
	ttl, err := s.store.TTL(key)
	if err != nil {
		return err
	}
	fmt.Fprintf(conn, "TTL: %d seconds\n", int(ttl.Seconds()))
	return nil
}
