package response

import "github.com/codecrafters-io/http-server-starter-go/app/encoding"

type Body struct {
	Content         string
	ContentType     string
	ContentLength   int
	ContentEncoding string
}

func NewBody(content string, contentType string) Body {
	return Body{
		Content:       content,
		ContentType:   contentType,
		ContentLength: len(content),
	}
}

func NewEncodedBody(content string, contentType string, encoder encoding.Encoder) (Body, error) {
	encoded, err := encoder.Encode(content)
	if err != nil {
		return Body{}, err
	}
	length := len(encoded)
	return Body{
		Content:         encoded,
		ContentType:     contentType,
		ContentLength:   length,
		ContentEncoding: encoder.Name,
	}, nil
}
