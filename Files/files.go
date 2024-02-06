package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateDirWithChildFiles() {
	// create the dir
	dirPath := "./vcs"
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("there was aproble")
		log.Fatal(err)
	}
	fileNames := []string{"config.txt", "index.txt"}
	// create the files and add them to the parent dir
	for _, fileName := range fileNames {
		filepath := filepath.Join(dirPath, fileName)
		file, err := os.Create(filepath)
		if err != nil {
			fmt.Println("error create the file.", err)
			return
		}
		defer file.Close()
	}
}
