package files

import (
	"os"
)

func CreateFile(filePath string) (*os.File, error) {
	var err error
	createdFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return createdFile, nil
}
