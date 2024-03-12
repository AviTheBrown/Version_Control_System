package files

import (
	datatypes "Version_Control_System/DataTypes"
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
		fmt.Println("there was a problem. ")
		log.Fatal(err)
	}

	commitDir := "commits"
	commitDirPath := filepath.Join(dirPath, commitDir)
	fmt.Println(commitDirPath)
	err = os.Mkdir(commitDirPath, 0755)
	if err != nil {
		fmt.Println("Error creating commits directory.")
		return
	}

	fileNames := []string{"config.txt", "index.txt"}
	// create the files and add them to the parent dir
	for _, fileName := range fileNames {
		filepath := filepath.Join(dirPath, fileName)
		file, err := os.Create(filepath)
		if err != nil {
			fmt.Println("error creating the file.", err)
			return
		}
		defer file.Close()
	}
}

func CreatHashDir(commitMsg string, hashString string, user datatypes.User) {
	var err error
	commitDir := filepath.Join(".", "vcs", "commits")
	commitsPath := filepath.Join(commitDir, hashString)

	_, err = os.Stat(commitsPath)

	if err == nil {
		fmt.Println("commits already committed")
	} else if os.IsNotExist(err) {
		err := os.MkdirAll(commitsPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Changes are committed")

	} else {
		log.Fatal(err)
	}
	for _, file := range user.FileInfo.FileNames {
		sourceIndexedFilePath := datatypes.INDEXFILEPATH
		// sourceConfigFilePath := datatypes.CONFIGFILEPATH
		copiedIndexedFilePath := filepath.Join(commitsPath, file)
		_, err := os.Stat(copiedIndexedFilePath)

		if err == nil {
			fmt.Println("file already created in the corresponding commit dir")
		} else if os.IsNotExist(err) {

			data, err := os.ReadFile(sourceIndexedFilePath)
			if err != nil {
				fmt.Println("Error reading file")
				return
			}

			err = os.WriteFile(copiedIndexedFilePath, data, 0644)
			if err != nil {
				fmt.Println("had a problem createing file in commits dir.")
				log.Fatal(err)
			}
		}
	}
	fmt.Println("done")
	fmt.Println("Files after:")
	fmt.Println(user.FileInfo.FileNames)
	fmt.Printf("the hash is: %v", hashString)
}
