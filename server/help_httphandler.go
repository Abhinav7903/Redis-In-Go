package server

import "net/http"

// handlerHelp returns an HTTP handler that provides a list of available commands and their usage.
// handlerHelp returns an HTTP handler that provides a list of available commands and their usage with curl examples.
func (s *Server) handlerHelp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		helpText := `Available commands and their usage:

1. SET key value1 value2 ...
   - Stores one or more values under the specified key.
   - Example: 
     - Command: SET mykey value1 value2 value3
     - Curl: 
       curl -X POST http://localhost:1234/set/mykey -d '["value1", "value2", "value3"]'

2. GET key
   - Retrieves all values associated with the specified key.
   - Example: 
     - Command: GET mykey
     - Curl:
       curl -X GET http://localhost:1234/get/mykey

3. DELETE key
   - Removes the specified key and its associated values from the store.
   - Example: 
     - Command: DELETE mykey
     - Curl:
       curl -X DELETE http://localhost:1234/delete/mykey

4. EXISTS key
   - Checks if the specified key exists in the store.
   - Example:
     - Command: EXISTS mykey
     - Curl:
       curl -X GET http://localhost:1234/exists/mykey

5. EXPIRE key ttl_in_seconds
   - Sets a time-to-live (TTL) for the specified key in seconds.
   - After the TTL expires, the key will be automatically removed.
   - Example: 
     - Command: EXPIRE mykey 60
     - Curl:
       curl -X POST http://localhost:1234/expire/mykey -d '{"ttl": 60}'

6. TTL key
   - Retrieves the remaining time-to-live (TTL) for the specified key.
   - Example: 
     - Command: TTL mykey
     - Curl:
       curl -X GET http://localhost:1234/ttl/mykey

7. RAND key offset
   - Retrieves values from the specified key starting from the given offset.
   - Example:
     - Command: RAND mykey 2
     - Curl:
       curl -X GET http://localhost:1234/rand/mykey/2

8. SETUQ key value1 value2 ...
   - Stores unique values under the specified key, avoiding duplicates.
   - Example: 
     - Command: SETUQ mykey value1 value2 value2 (value2 will be stored only once)
     - Curl:
       curl -X POST http://localhost:1234/setuq/mykey -d '["value1", "value2", "value3"]'

9. REMOVE key value
   - Removes a specific value from the specified key.
   - Example:
     - Command: REMOVE mykey value1
     - Curl:
       curl -X POST http://localhost:1234/remove/mykey -d '["value1"]'

10. GETUQ key
    - Retrieves all unique values associated with the specified key.
    - Example: 
      - Command: GETUQ mykey
      - Curl:
        curl -X GET http://localhost:1234/getuq/mykey

11. GETKEY value
    - Retrieves the key associated with the specified value.
    - Example: 
      - Command: GETKEY value1
      - Curl:
        curl -X GET http://localhost:1234/getkey/value1

13. HELP
    - Displays this help message.
    - Example:
      - Command: HELP
      - Curl:
        curl -X GET http://localhost:1234/help

For any issues or questions, please help yourself.
`

		s.respond(w, ResponseMsg{Message: "success", Data: helpText}, http.StatusOK, nil)

	}
}
