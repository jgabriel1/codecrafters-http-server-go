package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

const (
	OK_RESPONSE          = "HTTP/1.1 200 OK\r\n\r\n"
	BAD_REQUEST_RESPONSE = "HTTP/1.1 404 Not Found\r\n\r\n"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		exitWithMessage("Failed to bind to port 4221")
	}
	conn, err := l.Accept()
	if err != nil {
		exitWithMessage("Error accepting connection: ", err.Error())
	}
	defer conn.Close()
	lines, err := ReadAllLines(bufio.NewReader(conn))
	fmt.Println(lines)
	if err != nil {
		exitWithMessage("Error reading message: ", err.Error())
	}
	if len(lines) < 2 {
		exitWithMessage("Invalid message.")
	}
	path := lines[1]
	switch path {
	case "/":
		conn.Write([]byte(OK_RESPONSE))
	default:
		conn.Write([]byte(BAD_REQUEST_RESPONSE))
	}
}

func exitWithMessage(message ...any) {
	fmt.Println(message...)
	os.Exit(1)
}

func ReadAllLines(reader *bufio.Reader) ([]string, error) {
	lines := make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return lines, nil
			}
			return nil, err
		}
		lines = append(lines, string(line))
		fmt.Println(lines)
	}
}
