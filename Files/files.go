package files

import (
	"os"
)

func CreateFile(filePath string) (*os.File, error) {
	createdFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return createdFile, nil
}

func OpenFile(filepath string) (*os.File, error) {
	openedFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return openedFile, nil
}
