package request

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    string
}

func FromReader(reader io.Reader) (*Request, error) {
	lineReader := bufio.NewReader(reader)
	lines, err := readRequestHeadersLines(lineReader)
	if err != nil {
		return nil, err
	}
	method, path := parseFirstLine(lines[0])
	headers := parseHeaders(lines[1:])
	body := ""
	if lengthStr, ok := headers["content-length"]; ok {
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return nil, errors.New("invalid headers")
		}
		body, err = readRequestBody(lineReader, length)
		if err != nil {
			return nil, err
		}
	}
	return &Request{
		Method:  method,
		Path:    path,
		Headers: headers,
		Body:    body,
	}, nil
}

func parseFirstLine(line string) (string, string) {
	split := strings.Split(line, " ")
	return split[0], split[1]
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

func readRequestHeadersLines(reader *bufio.Reader) ([]string, error) {
	lines := make([]string, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		if len(line) == 0 { // zero size line to stop reading headers
			return lines, nil
		}
		lines = append(lines, string(line))
	}
}

func readRequestBody(reader io.Reader, length int) (string, error) {
	buf := make([]byte, length)
	_, err := reader.Read(buf)
	if err != nil {
		return "", err
	}
	trimmed := bytes.Trim(buf, "\x00")
	return string(trimmed), nil
}
