package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/config"
	"github.com/codecrafters-io/http-server-starter-go/app/filesystem"
	"github.com/codecrafters-io/http-server-starter-go/app/request"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
)

func main() {
	cfg := config.Parse()
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		exitWithMessage("Failed to bind to port 4221")
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleConnection(conn, cfg)
	}
}

func handleConnection(conn net.Conn, cfg config.Config) error {
	defer conn.Close()
	req, err := request.FromReader(conn)
	if err != nil {
		exitWithMessage("Error reading message: ", err.Error())
	}
	splitPath := strings.Split(req.Path, "/")
	switch splitPath[1] {
	case "":
		{
			res := response.NewText(response.StatusOk, "")
			_, err := conn.Write(response.Encode(res))
			return err
		}
	case "files":
		{
			fileName := splitPath[2]
			if req.Method == "GET" {
				content, err := filesystem.ReadFile(cfg, fileName)
				if err != nil {
					res := response.NewText(response.StatusNotFound, "")
					_, err = conn.Write(response.Encode(res))
					return err
				}
				res := response.New(
					response.StatusOk,
					response.NewBody(string(content), "application/octet-stream"))
				_, err = conn.Write(response.Encode(res))
				return err
			} else {
				if err = filesystem.WriteToFile(cfg, req.Body, fileName); err != nil {
					res := response.NewText(response.StatusNotFound, "")
					_, err := conn.Write(response.Encode(res))
					return err
				}
				res := response.NewText(response.StatusCreated, "")
				_, err := conn.Write(response.Encode(res))
				return err
			}
		}
	case "user-agent":
		{
			res := response.NewText(response.StatusOk, req.Headers["user-agent"])
			_, err := conn.Write(response.Encode(res))
			return err
		}
	case "echo":
		{
			message := splitPath[2]
			res := response.NewText(response.StatusOk, message)
			_, err := conn.Write(response.Encode(res))
			return err
		}
	default:
		{
			res := response.NewText(response.StatusNotFound, "")
			_, err := conn.Write(response.Encode(res))
			return err
		}
	}
}

func exitWithMessage(message ...any) {
	fmt.Println(message...)
	os.Exit(1)
}
