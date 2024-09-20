package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <key>")
		return
	}

	key := os.Args[1]

	conn, err := net.Dial("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Generate values from 1 to 2000
	var values []string
	for i := 1; i <= 100; i++ {
		values = append(values, strconv.Itoa(i))
	}

	command := fmt.Sprintf("SET %s %s\n", key, strings.Join(values, " "))
	_, err = conn.Write([]byte(command))
	if err != nil {
		log.Fatalf("Failed to send command: %v", err)
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	fmt.Println("Server response:", response)
}
