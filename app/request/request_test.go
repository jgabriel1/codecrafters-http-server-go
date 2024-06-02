package request

import (
	"strings"
	"testing"
)

const ECHO_REQUEST = "GET /echo/abc HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n"

func TestReadsPathCorrectly(t *testing.T) {
	req, _ := FromReader(strings.NewReader(ECHO_REQUEST))
	if req.Path != "/echo/abc" {
		t.Error("does not read path correctly")
	}
}
