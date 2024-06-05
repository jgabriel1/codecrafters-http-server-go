package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodesOkResponseCorrectly(t *testing.T) {
	res := NewText(StatusOk, "foo")
	encoded := Encode(res)
	expected := []byte(
		"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 3\r\n\r\nfoo",
	)
	assert.EqualValues(t, expected, encoded, "does not encode correctly")
}

func TestEncodesNotFoundResponseCorrectly(t *testing.T) {
	res := NewText(StatusNotFound, "")
	encoded := Encode(res)
	expected := []byte(
		"HTTP/1.1 404 Not Found\r\n\r\n",
	)
	assert.EqualValues(t, expected, encoded, "does not encode correctly")
}

func TestAddsEncodingHeader(t *testing.T) {
	res := NewText(StatusOk, "foo")
	encoded := EncodeWith(res, "gzip")
	expected := []byte(
		"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 3\r\nContent-Encoding: gzip\r\n\r\nfoo",
	)
	assert.EqualValues(t, expected, encoded, "does not add encoding header")
}
