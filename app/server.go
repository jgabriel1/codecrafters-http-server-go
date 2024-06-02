package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/request"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
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
	req, err := request.FromReader(conn)
	if err != nil {
		exitWithMessage("Error reading message: ", err.Error())
	}
	matchPath(conn, req.Path)
}

func matchPath(conn net.Conn, path string) error {
	splitPath := strings.Split(path, "/")
	switch splitPath[1] {
	case "echo":
		{
			res := response.New(response.StatusOk, splitPath[2])
			_, err := conn.Write(response.Encode(res))
			return err
		}
	default:
		{
			res := response.New(response.StatusNotFound, "")
			_, err := conn.Write(response.Encode(res))
			return err
		}
	}
}

func exitWithMessage(message ...any) {
	fmt.Println(message...)
	os.Exit(1)
}
