package main

import (
	"fmt"
	"net"

	"github.com/bluesberries/my_redis_server/resp"
)

func main() {
	setupTCPConnection()
}

func setupTCPConnection() {
	listener, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Server is listening on port 6379")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		//Process and use the data (here, we'll just print it)
		decoded_response, err := resp.Deserialize(buffer[:n])
		fmt.Printf("Received:\n %s\n", string(decoded_response))

		//data := []byte("$4\r\nPONG\r\n")
		data := []byte("+PONG\r\n")
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		return
	}

}
