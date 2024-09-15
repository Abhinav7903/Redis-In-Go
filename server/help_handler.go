package server

import (
	"fmt"
	"net"
)

func (s *Server) handleHelp(conn net.Conn) error {
	helpText := `Available commands and their usage:

1. SET key value1 value2 ...
   - Stores one or more values under the specified key.
   - Example: SET mykey value1 value2 value3

2. GET key
   - Retrieves all values associated with the specified key.
   - Example: GET mykey

3. DELETE key
   - Removes the specified key and its associated values from the store.
   - Example: DELETE mykey

4. EXISTS key
   - Checks if the specified key exists in the store.
   - Returns 1 if the key exists, 0 otherwise.
   - Example: EXISTS mykey

5. EXPIRE key ttl_in_seconds
   - Sets a time-to-live (TTL) for the specified key in seconds.
   - After the TTL expires, the key will be automatically removed.
   - Example: EXPIRE mykey 60 (sets TTL of 60 seconds)

6. TTL key
   - Retrieves the remaining time-to-live (TTL) for the specified key.
   - Returns the TTL in seconds.
   - Example: TTL mykey

7. RAND key offset
   - Retrieves values from the specified key starting from the given offset.
   - Example: RAND mykey 2 (retrieves values from index 2 onward)

8. SETUQ key value1 value2 ...
   - Stores unique values under the specified key, avoiding duplicates.
   - Example: SETUQ mykey value1 value2 value2 (value2 will be stored only once)

9. REMOVE key value
   - Removes a specific value from the specified key.
   - Example: REMOVE mykey value1

10. GETUQ key
    - Retrieves all unique values associated with the specified key.
    - Example: GETUQ mykey

11. GETKEY value
    - Retrieves the key associated with the specified value.
    - Example: GETKEY value1

12. EXIT
    - Closes the connection and exits the session.

13. HELP
    - Displays this help message.

For any issues or questions, please help yourself.
`
	fmt.Fprint(conn, helpText)
	return nil
}
