package encoding

import (
	"bytes"
	"compress/gzip"
)

type Encoder struct {
	Name   string
	Encode func(payload string) (string, error)
}

var encoders = map[string]Encoder{
	"gzip": {
		Name: "gzip",
		Encode: func(payload string) (string, error) {
			var buf bytes.Buffer
			gzipWriter := gzip.NewWriter(&buf)
			defer gzipWriter.Close()
			_, err := gzipWriter.Write([]byte(payload))
			if err != nil {
				return "", err
			}
			if err := gzipWriter.Close(); err != nil {
				return "", err
			}
			return buf.String(), nil
		},
	},
}

func FindValidEncoder(encodings []string) *Encoder {
	for _, encoding := range encodings {
		if encoder, ok := encoders[encoding]; ok {
			return &encoder
		}
	}
	return nil
}
