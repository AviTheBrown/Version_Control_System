package files

import (
	"fmt"
	"os"
)

func CreateFile(filePath string) (*os.File, error) {
	createdFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to create file")
		return nil, err
	}
	return createdFile, nil
}

func ReadFile(filepath string) (*os.File, error) {
	openedFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println("failed to open read only file.")
		return nil, err
	}

	return openedFile, nil
}
