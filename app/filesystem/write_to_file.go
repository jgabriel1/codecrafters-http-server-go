package filesystem

import (
	"os"
	"path"

	"github.com/codecrafters-io/http-server-starter-go/app/config"
)

func WriteToFile(cfg config.Config, content string, name string) error {
	filePath := path.Join(cfg.FilesDirectory, name)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
