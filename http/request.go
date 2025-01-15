package http

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Request struct {
	Method  string
	URI     string
	Headers map[string]string
	Body    string
}

func parseRequest(conn net.Conn) (*Request, error) {
	reader := bufio.NewReader(conn)
	line, _, err := reader.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("failed to read request line: %v", err)
	}

	parts := strings.Fields(string(line))
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid request line")
	}

	req := &Request{
		Method:  parts[0],
		URI:     parts[1],
		Headers: make(map[string]string),
	}

	for {
		line, _, err = reader.ReadLine()
		if err != nil || len(line) == 0 {
			break
		}
		headerParts := strings.SplitN(string(line), ":", 2)
		if len(headerParts) == 2 {
			req.Headers[strings.TrimSpace(headerParts[0])] = strings.TrimSpace(headerParts[1])
		}
	}

	if req.Method == "POST" {
		body, _ := reader.ReadString('\n')
		req.Body = body
	}

	return req, nil
}
