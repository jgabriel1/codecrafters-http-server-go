package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/config"
	"github.com/codecrafters-io/http-server-starter-go/app/encoding"
	"github.com/codecrafters-io/http-server-starter-go/app/filesystem"
	"github.com/codecrafters-io/http-server-starter-go/app/request"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
	"github.com/codecrafters-io/http-server-starter-go/app/router"
)

func main() {
	cfg := config.Parse()
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		exitWithMessage("Failed to bind to port 4221")
	}
	appRouter := buildRouter()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		req, err := request.FromReader(conn)
		if err != nil {
			conn.Write(response.Encode(*response.DefaultInternalServerError()))
			conn.Close()
			continue
		}
		go handle(appRouter, conn, *req, cfg)
	}
}

func buildRouter() *router.Router {
	r := router.New()
	r.AddRoute("/", "GET",
		func(req request.Request, cfg config.Config) *response.Response {
			res := response.NewText(response.StatusOk, "")
			return &res
		})
	r.AddRoute("/files/:filename", "GET",
		func(req request.Request, cfg config.Config) *response.Response {
			fileName := strings.Split(req.Path, "/")[2]
			content, err := filesystem.ReadFile(cfg, fileName)
			if err != nil {
				res := response.NewText(response.StatusNotFound, "")
				return &res
			}
			res := response.New(
				response.StatusOk,
				response.NewBody(string(content), "application/octet-stream"))
			return &res
		})
	r.AddRoute("/files/:filename", "POST",
		func(req request.Request, cfg config.Config) *response.Response {
			fileName := strings.Split(req.Path, "/")[2]
			if err := filesystem.WriteToFile(cfg, req.Body, fileName); err != nil {
				res := response.NewText(response.StatusNotFound, "")
				return &res
			}
			res := response.NewText(response.StatusCreated, "")
			return &res
		})
	r.AddRoute("/user-agent", "GET",
		func(req request.Request, cfg config.Config) *response.Response {
			res := response.NewText(response.StatusOk, req.Headers["user-agent"])
			return &res
		})
	r.AddRoute("/echo/:message", "GET",
		func(req request.Request, cfg config.Config) *response.Response {
			message := strings.Split(req.Path, "/")[2]
			if encodings, ok := req.Headers["accept-encoding"]; ok {
				encodingsList := regexp.MustCompile(`\s*,\s*`).Split(encodings, -1)
				encoder := encoding.FindValidEncoder(encodingsList)
				if encoder != nil {
					body, err := response.NewEncodedBody(message, "text/plain", *encoder)
					if err == nil {
						res := response.New(response.StatusOk, body)
						return &res
					}
				}
			}
			res := response.NewText(response.StatusOk, message)
			return &res
		})
	return r
}

func handle(r *router.Router, conn net.Conn, req request.Request, cfg config.Config) {
	defer conn.Close()
	res := r.Handle(req, cfg)
	conn.Write(response.Encode(*res))
}

func exitWithMessage(message ...any) {
	fmt.Println(message...)
	os.Exit(1)
}
