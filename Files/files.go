package files

import (
	"fmt"
	"log"
	"os"
)

func CreateDirWithChildFiles() {
	// create the dir
	dirPath := "./vcs"
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	fileNames := []string{"config.tx", "index.tx"}

	// create the files and add them to the parent dir
	for _, fileName := range fileNames {
		filepath := dirPath + "/" + fileName
		file, err := os.Create(filepath)
		if err != nil {
			fmt.Println("error create the file.", err)
			return
		}
		defer file.Close()
	}
}
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
