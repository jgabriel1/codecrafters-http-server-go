package response

import (
	"bytes"
	"testing"
)

func TestEncodesOkResponseCorrectly(t *testing.T) {
	res := New(StatusOk, "foo")
	encoded := Encode(res)
	expected := []byte(
		"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 3\r\n\r\nfoo",
	)
	if !bytes.Equal(encoded, expected) {
		t.Error("does not encode correctly")
	}
}

func TestEncodesNotFoundResponseCorrectly(t *testing.T) {
	res := New(StatusNotFound, "")
	encoded := Encode(res)
	expected := []byte(
		"HTTP/1.1 404 Not Found\r\n\r\n",
	)
	if !bytes.Equal(encoded, expected) {
		t.Error("does not encode correctly")
	}
}
