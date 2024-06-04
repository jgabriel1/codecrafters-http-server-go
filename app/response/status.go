package response

type Status uint8

const (
	StatusOk = iota
	StatusCreated
	StatusNotFound
)

func (s Status) ToEncoded() string {
	switch s {
	case StatusOk:
		return "HTTP/1.1 200 OK"
	case StatusCreated:
		return "HTTP/1.1 201 Created"
	case StatusNotFound:
		return "HTTP/1.1 404 Not Found"
	default:
		return ""
	}
}
