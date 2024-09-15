package server

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func (s *Server) handleRand(conn net.Conn, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: RAND key offset")
	}
	key, offsetStr := args[0], args[1]
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return fmt.Errorf("invalid offset value")
	}
	values, err := s.store.RandomValues(key, offset)
	if err != nil {
		return err
	}
	fmt.Fprint(conn, "Values: "+strings.Join(values, ", ")+"\n")
	return nil
}
