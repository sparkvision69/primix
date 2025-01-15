package http

import (
	"fmt"
	"net"
)

func StartServer(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	defer listener.Close()

	fmt.Printf("Server started on port %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	req, err := parseRequest(conn)
	if err != nil {
		fmt.Println("Error parsing request:", err)
		return
	}

	response := &Response{
		StatusCode: 200,
		Body:       "Hello from custom HTTP server!",
	}
	sendResponse(conn, response)

	fmt.Printf("Received %s request for %s\n", req.Method, req.URI)
}
