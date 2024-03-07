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

func creatHashDir(hash string) {
	startingDir, _ := os.Getwd()
	fmt.Println(startingDir)
	// defer os.Chdir(startingDir)
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println("There was a problem retreiving the PWD")
	// 	return
	// }

	// fmt.Printf("Current directory is: %v", startingDir)
	// commitDir := ".vcs/commits"
	// err = os.Chdir(commitDir)
	// if err != nil {
	// 	fmt.Println("There was a problem changing directories.")
	// 	return
	// }
	// _, err = os.Stat(hash)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		err = os.Mkdir(hash, 0755)
	// 	}
	// }

}
