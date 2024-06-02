package main

import (
	"bufio"
	"net"
	"strings"
)

type Request struct {
	Path string
}

func NewRequest(conn net.Conn) (*Request, error) {
	lines, err := readRequestLines(conn)
	if err != nil {
		return nil, err
	}
	path := strings.Split(lines[0], " ")[1]
	return &Request{
		Path: path,
	}, nil
}

func readRequestLines(c net.Conn) ([]string, error) {
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
