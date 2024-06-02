package response

type Body struct {
	Content       string
	ContentType   string
	ContentLength int
}

func NewBody(content string, contentType string) Body {
	length := len(content)
	return Body{
		Content:       content,
		ContentType:   contentType,
		ContentLength: length,
	}
}
