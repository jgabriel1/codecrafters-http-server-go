package request

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const EXAMPLE_ECHO_REQUEST = "GET /echo/abc HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n"

func TestReadsPath(t *testing.T) {
	req, err := FromReader(strings.NewReader(EXAMPLE_ECHO_REQUEST))
	assert.Nil(t, err, "errors reading request")
	assert.Equal(t, req.Path, "/echo/abc", "does not read path correctly")
}

func TestReadsHeaders(t *testing.T) {
	req, err := FromReader(strings.NewReader(EXAMPLE_ECHO_REQUEST))
	assert.Nil(t, err, "errors reading request")
	assert.EqualValues(t, map[string]string{
		"host":       "localhost:4221",
		"user-agent": "curl/7.64.1",
		"accept":     "*/*",
	}, req.Headers, "does not read headers correctly")
}

func TestReadsEmptyBody(t *testing.T) {
	req, err := FromReader(strings.NewReader(EXAMPLE_ECHO_REQUEST))
	assert.Nil(t, err, "does not read empty request")
	assert.Empty(t, req.Body, "request body is not empty")
}

const EXAMPLE_BODY_REQUEST = "POST /path HTTP/1.1\r\nHost: example.com\r\nContent-Type: text/plain\r\nContent-Length: 3\r\n\r\nfoo"

func TestReadsBody(t *testing.T) {
	req, err := FromReader(strings.NewReader(EXAMPLE_BODY_REQUEST))
	assert.Nil(t, err, "does not read request")
	assert.Equal(t, "foo", req.Body, "reads body incorrectly")
}

func TestReadsOnlyContentLenghtBytes(t *testing.T) {
	reqString := EXAMPLE_BODY_REQUEST + "bar"
	req, err := FromReader(strings.NewReader(reqString))
	assert.Nil(t, err, "errors reading request")
	assert.Equal(t, "foo", req.Body, "does not read only the specified number of bytes")
}

const EXAMPLE_BODY_REQUEST_WITH_WRONG_LENGTH = "POST /path HTTP/1.1\r\nHost: example.com\r\nContent-Type: text/plain\r\nContent-Length: 10\r\n\r\nfoo"

func TestNoErrorsWhenContentLengthGreaterThanAvailableContent(t *testing.T) {
	req, err := FromReader(strings.NewReader(EXAMPLE_BODY_REQUEST_WITH_WRONG_LENGTH))
	assert.Nil(t, err, "should error reading request")
	assert.Equal(t, "foo", req.Body, "does not read only the available bytes")
}
