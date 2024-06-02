package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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
	lines, err := ReadAllLines(conn)
	if err != nil {
		exitWithMessage("Error reading message: ", err.Error())
	}
	if len(lines) < 2 {
		exitWithMessage("Invalid message.")
	}
	path := strings.Split(lines[0], " ")[1]
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

func ReadAllLines(c net.Conn) ([]string, error) {
	reader := bufio.NewReader(c)
	lines := make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		if len(line) == 0 { // zero size chunk to stop reading
			return lines, nil
		}
		lines = append(lines, string(line))
	}
}
