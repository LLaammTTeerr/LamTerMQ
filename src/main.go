package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	conn.Read(buffer)
	resp := buffer
	conn.Write(resp)
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Failed to start server: %v\nMake sure port 8080 is not in use\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP Server listening on :8080")
	fmt.Println("You can connect using: nc localhost 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
