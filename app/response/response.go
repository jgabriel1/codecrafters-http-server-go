package response

import (
	"fmt"
	"strings"
)

type Response struct {
	body   Body
	status Status
}

func New(status Status, body Body) Response {
	return Response{
		body:   body,
		status: status,
	}
}

func NewText(status Status, text string) Response {
	body := NewBody(text, "text/plain")
	return Response{
		body:   body,
		status: status,
	}
}

func Encode(res Response) []byte {
	content := []string{}
	content = append(content, res.status.ToEncoded())
	if res.body.ContentLength > 0 {
		content = append(content,
			fmt.Sprintf("Content-Type: %s", res.body.ContentType),
			fmt.Sprintf("Content-Length: %d", res.body.ContentLength),
		)
	}
	content = append(content,
		"", // empty string to signal the end of the headers
		res.body.Content,
	)
	return []byte(strings.Join(content, "\r\n"))
}
