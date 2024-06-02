package request

import (
	"bufio"
	"io"
	"strings"
)

type Request struct {
	Path string
}

func FromReader(reader io.Reader) (*Request, error) {
	lines, err := readRequestLines(reader)
	if err != nil {
		return nil, err
	}
	path := strings.Split(lines[0], " ")[1]
	return &Request{
		Path: path,
	}, nil
}

func readRequestLines(reader io.Reader) ([]string, error) {
	r := bufio.NewReader(reader)
	lines := make([]string, 0)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			return nil, err
		}
		if len(line) == 0 { // zero size chunk to stop reading
			return lines, nil
		}
		lines = append(lines, string(line))
	}
}
