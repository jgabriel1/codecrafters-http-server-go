package request

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type Request struct {
	Path    string
	Headers map[string]string
	Body    string
}

func FromReader(reader io.Reader) (*Request, error) {
	lines, err := readRequestHeadersLines(reader)
	if err != nil {
		return nil, err
	}
	_, path, _ := parseFirstLine(lines[0])
	headers := parseHeaders(lines[1:])
	body := ""
	if length, ok := headers["content-length"]; ok {
		body, err = readRequestBody(reader, length)
		if err != nil {
			return nil, err
		}
	}
	return &Request{
		Path:    path,
		Headers: headers,
		Body:    body,
	}, nil
}

func parseFirstLine(line string) (string, string, string) {
	split := strings.Split(line, " ")
	return split[0], split[1], split[2]
}

func parseHeaders(headersLines []string) map[string]string {
	headers := map[string]string{}
	for _, line := range headersLines {
		splitHeader := strings.Split(line, ": ")
		key := strings.ToLower(splitHeader[0])
		headers[key] = splitHeader[1]
	}
	return headers
}

func readRequestHeadersLines(reader io.Reader) ([]string, error) {
	r := bufio.NewReader(reader)
	lines := make([]string, 0)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			return nil, err
		}
		if len(line) == 0 { // zero size line to stop reading headers
			return lines, nil
		}
		lines = append(lines, string(line))
	}
}

func readRequestBody(reader io.Reader, length string) (string, error) {
	return "", errors.New("not implemented")
}
