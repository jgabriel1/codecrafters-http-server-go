package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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
	req, err := NewRequest(conn)
	if err != nil {
		exitWithMessage("Error reading message: ", err.Error())
	}
	matchPath(conn, req.Path)
}

func matchPath(conn net.Conn, path string) error {
	splitPath := strings.Split(path, "/")
	switch splitPath[0] {
	case "echo":
		{
			res, _ := NewResponse(splitPath[1])
			return res.Write(conn)
		}
	default:
		return nil
	}
}

func exitWithMessage(message ...any) {
	fmt.Println(message...)
	os.Exit(1)
}
