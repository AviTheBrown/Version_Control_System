package files

import (
	datatypes "Version_Control_System/DataTypes"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func AllFilesCreated() bool {
	fileNames := []string{"config.txt", "index.txt", "log.txt"}
	for _, fileName := range fileNames {
		filePath := filepath.Join("vcs", fileName)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return false
		}
	}
	return true
}
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
	fmt.Println("commits dir path files.go")
	fmt.Println(commitDirPath)
	err = os.Mkdir(commitDirPath, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Error creating commits directory.")
		return
	}

	fileNames := []string{"config.txt", "index.txt", "log.txt"}
	// create the files and add them to the parent dir
	for _, fileName := range fileNames {
		filepath := filepath.Join(dirPath, fileName)
		_, err := os.Stat(filepath)
		if os.IsNotExist(err) {
			file, err := os.Create(filepath)
			if err != nil {
				fmt.Println("error creating the file.", err)
				return
			}
			defer file.Close()
		}
	}
}

func CreatHashDir(commitMsg string, hashString string, user datatypes.User) {
	var err error
	// vcs/commits
	commitDir := filepath.Join(".", "vcs", "commits")
	// vcs/commits/<hashString>
	commitHashDir := filepath.Join(commitDir, hashString)

	_, err = os.Stat(commitHashDir)

	if err == nil {
		fmt.Println("Commits already committed")
		return
	} else if os.IsNotExist(err) {
		err := os.MkdirAll(commitHashDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Commit Hash Dir Created.")

	} else {
		log.Fatal(err)
	}
	// _, indexFile := filepath.Split(datatypes.INDEXFILEPATH)
	for _, file := range user.FileInfo.FileNames {
		// vcs/commits/<hashDir>/file.txt
		completedHashDirFilePath := filepath.Join(commitHashDir, file)
		_, err := os.Stat(completedHashDirFilePath)

		if err == nil {
			fmt.Println("file already created in the corresponding commit dir")
		} else {
			err := os.MkdirAll(completedHashDirFilePath, 0755)
			if err != nil {
				fmt.Println("There was a problem creating the completed path.")
			}
		}
		fmt.Printf("%s: was added to the hashDir\n", file)

	}
	fmt.Println("done")
	fmt.Println("Files after:")
	fmt.Println(user.FileInfo.FileNames)
	fmt.Printf("the hash is: %v", hashString)
}
