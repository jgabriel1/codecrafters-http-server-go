package request

import (
	"maps"
	"strings"
	"testing"
)

const EXAMPLE_ECHO_REQUEST = "GET /echo/abc HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n"

func TestReadsPath(t *testing.T) {
	req, _ := FromReader(strings.NewReader(EXAMPLE_ECHO_REQUEST))
	if req.Path != "/echo/abc" {
		t.Error("does not read path correctly")
	}
}

func TestReadsHeaders(t *testing.T) {
	req, _ := FromReader(strings.NewReader(EXAMPLE_ECHO_REQUEST))
	passes := maps.Equal(req.Headers, map[string]string{
		"host":       "localhost:4221",
		"user-agent": "curl/7.64.1",
		"accept":     "*/*",
	})
	if !passes {
		t.Error("does not read headers correctly")
	}
}
