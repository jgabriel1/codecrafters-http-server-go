package response

func DefaultNotFound() *Response {
	res := NewText(StatusNotFound, "")
	return &res
}

func DefaultInternalServerError() *Response {
	res := NewText(StatusInternalServerError, "")
	return &res
}
