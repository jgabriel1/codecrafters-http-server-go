package response

import (
	"fmt"
	"strings"
)

type Response struct {
	body   string
	status Status
}

func New(status Status, body string) Response {
	return Response{
		body:   body,
		status: status,
	}
}

func Encode(res Response) []byte {
	content := []string{}
	content = append(content, res.status.ToEncoded())
	if len(res.body) > 0 {
		content = append(content,
			"Content-Type: text/plain",
			fmt.Sprintf("Content-Length: %d", len(res.body)),
		)
	}
	content = append(content,
		"", // empty string to signal the end of the headers
		res.body,
	)
	return []byte(strings.Join(content, "\r\n"))
}
