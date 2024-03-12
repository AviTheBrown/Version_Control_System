package datatypes

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type SVCS map[string]string

type User struct {
	UserName string
	FileInfo Files
}
type Files struct {
	FileNames []string
	CommitID  [32]byte
}

func CreateUser() *User {
	user := &User{
		FileInfo: Files{
			FileNames: []string{},
			CommitID:  [32]byte{},
		},
	}
	return user
}

const (
	CONFIGFILEPATH = "vcs/config.txt"
	INDEXFILEPATH  = "vcs/index.txt"
)

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
	appendToFile(INDEXFILEPATH, filename)
	u.FileInfo.FileNames = append(u.FileInfo.FileNames, filename)
	fmt.Println(formatOutput(filename, false))

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
