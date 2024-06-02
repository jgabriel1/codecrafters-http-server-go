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
	res, _ := NewResponse(strings.TrimLeft(req.Path, "/"))
	res.Write(conn)
}

func exitWithMessage(message ...any) {
	fmt.Println(message...)
	os.Exit(1)
}
