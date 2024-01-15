package main

import (
	"log"
	"os"
)

func CreateFile(filePath string) *os.File {
	openedFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return openedFile
}
