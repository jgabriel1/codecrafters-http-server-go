package filesystem

import (
	"errors"
	"os"
	"path"

	"github.com/codecrafters-io/http-server-starter-go/app/config"
)

func ReadFile(cfg config.Config, file string) ([]byte, error) {
	filePath := path.Join(cfg.FilesDirectory, file)
	if !fileExists(filePath) {
		return nil, errors.New("file does not exist")
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
