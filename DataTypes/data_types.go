package datatypes

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type SVCS map[string]string

type User struct {
	UserName string
	FileMeta []File
}
type File struct {
	FileName         string
	FileData         []byte
	FileCreationTime time.Time
	FileHash         string
}

const (
	CONFIGFILEPATH = "vcs/config.txt"
	INDEXFILEPATH  = "vcs/index.txt"
	LOGFILEPATH    = "vcs/log.txt"
)

func CreateUser() (*User, error) {
	userName, err := os.ReadFile(CONFIGFILEPATH)
	if err == nil {
		user := &User{
			UserName: string(userName),
			FileMeta: make([]File, 0),
		}
		return user, nil
	} else {
		return nil, fmt.Errorf("Error opening up CONFIG_FILE_PATH")
	}
}

func (u *User) AddFileToMeta(filename string, fileData []byte, fileHash string, fileCreationTime time.Time) {
	file := File{
		FileName:         filename,
		FileData:         fileData,
		FileHash:         fileHash,
		FileCreationTime: fileCreationTime,
	}
	u.FileMeta = append(u.FileMeta, file)
}
func (u *User) LoadUserName(filepath string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	u.UserName = string(content)
}
func (u *User) LoadTrackedFiles(filepath string) []string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading tracked files %v\n", err)
		return nil
	}
	files := strings.FieldsFunc(string(content), func(r rune) bool {
		return r == '\n' || r == '\r'
	})
	return files
}
func (u *User) AddAction(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Can't find '%s'.\n", filename)
		return
	}

	if u.isFileTracked(filename) {
		fmt.Println(formatOutput(filename, true))
		return
	}

	// Open the file in append mode
	indexFile, err := os.OpenFile(INDEXFILEPATH, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0655)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer indexFile.Close()

	// Write the filename to the file
	_, err = indexFile.WriteString(filename + "\n")
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file.")
	}

	fileInfo, err := os.Stat(filename)
	fileCreation := fileInfo.ModTime()

	sha256Hash := sha256.New()
	sha256Hash.Write(fileData)
	sha256HashValue := sha256Hash.Sum(nil)
	fileHash := hex.EncodeToString(sha256HashValue)

	u.AddFileToMeta(filename, fileData, fileHash, fileCreation)
	fmt.Println(formatOutput(filename, false))
	fmt.Println("done!")

}
func (u *User) isFileTracked(filename string) bool {
	trackedFiles := u.LoadTrackedFiles(INDEXFILEPATH)
	for _, trackedFile := range trackedFiles {
		if trackedFile == filename {
			return true
		}
	}
	return false
}

func appendToFile(filePath, content string) {
	switch {
	case filePath == CONFIGFILEPATH:
		err := os.WriteFile(filePath, []byte(content), os.ModePerm)
		if err != nil {
			fmt.Println("Error writing to file.")
		}
	case filePath == INDEXFILEPATH:

		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			fmt.Printf("Error writing to file %v\n", err)
		}
		defer file.Close()
		_, err = fmt.Fprint(file, content, "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func formatOutput(fileName string, isTracked bool) string {
	if isTracked {
		return fmt.Sprintf("'%s' is alread tracked.", fileName)
	}
	return fmt.Sprintf("The file '%s' is now tracked.", fileName)
}

func (u *User) ConfigAction(userName string) string {
	appendToFile(CONFIGFILEPATH, userName)
	return fmt.Sprintf("The username is %s.", userName)
}
