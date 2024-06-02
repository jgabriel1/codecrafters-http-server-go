package main

import (
	"fmt"
	"net"
	"strings"
)

type Response struct {
	body string
}

func NewResponse(body string) (*Response, error) {
	return &Response{
		body: body,
	}, nil
}

func (res *Response) encode() []byte {
	status := "HTTP/1.1 200 OK"
	contentTypeHeader := "Content-Type: text/plain"
	contentLengthHeader := fmt.Sprintf("Content-Length: %d", len(res.body))
	content := []string{
		status,
		contentTypeHeader,
		contentLengthHeader,
		"", // empty string to signal the end of the headers
		res.body,
	}
	return []byte(strings.Join(content, "\r\n"))
}

func (res *Response) Write(conn net.Conn) error {
	_, err := conn.Write(res.encode())
	return err
}
