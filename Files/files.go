package files

import (
	datatypes "Version_Control_System/DataTypes"
	"bufio"
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
func CreateVCSDirWithChildFiles() {
	// create the dir
	dirPath := "./vcs"
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("there was a problem. ")
		log.Fatal(err)
	}

	commitDir := "commits"
	commitDirPath := filepath.Join(dirPath, commitDir)
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

func CreateHashDir(commitMsg string, hashString string, user datatypes.User) {
	var err error
	// vcs/commits
	commitDir := filepath.Join(".", "vcs", "commits")
	// vcs/commits/<hashString>
	commitHashDir := filepath.Join(commitDir, hashString)

	_, err = os.Stat(commitHashDir)

	if os.IsNotExist(err) {
		err := os.MkdirAll(commitHashDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Commit Hash Dir Created.")

	}
	//creates the files that are in index.txt and add them to hashed commits dir
	for _, file := range user.FileMeta {
		// vcs/commits/<hashDir>/file.txt
		// file, _ := os.Open()
		completedHashDirFilePath := filepath.Join(commitHashDir, file.FileName)
		err := os.WriteFile(completedHashDirFilePath, file.FileData, 0655)
		if err != nil {
			fmt.Println("Error writting to file.")
			continue
		}
		_, err = os.Stat(completedHashDirFilePath)

		if err == nil {
			fmt.Println("file already created in the corresponding commit dir")
		} else {
			err := os.MkdirAll(completedHashDirFilePath, 0755)
			if err != nil {
				fmt.Println("There was a problem creating the completed path.")
			}
		}

	}
	fmt.Printf("the hash is: %v", hashString)
}
func LogAction(fileCommit string, author string, commitMessage string) {
	fmt.Println("in log file.")
	_, err := os.Stat(datatypes.LOGFILEPATH)
	if err != nil {
		fmt.Println("Error finding file")
		return
	}
	logFile, err := os.OpenFile(datatypes.LOGFILEPATH, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0655)
	_, err = os.ReadFile(datatypes.LOGFILEPATH)
	defer logFile.Close()

	writer := bufio.NewWriter(logFile)
	writer.WriteString("\n")
	_, err = writer.WriteString("commit " + fileCommit + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	_, err = writer.WriteString("Author: " + author + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	_, err = writer.WriteString(commitMessage + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("there was a problem flusig the writer.")
	}
}
