package http

import (
	"fmt"
	"net"
)

type Response struct {
	StatusCode int
	Body       string
}

func sendResponse(conn net.Conn, res *Response) {
	status := fmt.Sprintf("HTTP/1.1 %d OK\r\n", res.StatusCode)
	headers := "Content-Type: text/plain\r\n" 
	body := res.Body

	fmt.Fprintf(conn, "%s%s\r\n%s", status, headers, body)
}
