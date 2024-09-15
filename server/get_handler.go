package server

import (
	"fmt"
	"net"
)

func (s *Server) handleGet(conn net.Conn, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: GET key")
	}
	key := args[0]
	values, err := s.store.Get(key)
	if err != nil {
		return err
	}

	if len(values) == 0 {
		fmt.Fprint(conn, "No values found for key: "+key+"\n")
		return nil
	}

	// Numbered list of values
	for i, value := range values {
		fmt.Fprintf(conn, "%d: %s\n", i+1, value)
	}
	return nil
}

func (s *Server) handleGetUnique(conn net.Conn, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: GETUQ key")
	}
	key := args[0]
	values, err := s.store.GetUnique(key)
	if err != nil {
		return err
	}

	if len(values) == 0 {
		fmt.Fprint(conn, "No unique values found for key: "+key+"\n")
		return nil
	}

	// Numbered list of unique values
	for i, value := range values {
		fmt.Fprintf(conn, "%d: %s\n", i+1, value)
	}
	return nil
}

func (s *Server) handleGetKey(conn net.Conn, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: GETKEY value")
	}
	value := args[0]
	keys, err := s.store.GetKeyFromValue(value)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		fmt.Fprint(conn, "value not found\n")
	} else {
		for _, key := range keys {
			fmt.Fprintf(conn, "Key: %s\n", key)
		}
	}

	return nil
}

// TODO fix the getkey and delete functions
