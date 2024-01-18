package datatypes

import (
	files "Version_Control_System/Files"
	"fmt"
	"os"
)

type SVCS map[string]string

type FileInfo struct {
	Files     []*os.File
	FileNames []string
}
type User struct {
	UserName string
	FileInfo
}

func CreateUser() *User {
	return &User{}
}
func (u *User) ConfigAction(userName string) string {
	if u.UserName == userName {
		return fmt.Sprintf("The username is %s", u.UserName)
	}
	u.UserName = userName
	return fmt.Sprintf("The username is %s", u.UserName)
}

func containsFile(files []string, target string) bool {
	for _, element := range files {
		if element == target {
			return true
		}
	}
	return false
}
func (u *User) AddAction(fileName string) string {

	var err error
	// create the file
	openedFile, err := files.CreateFile(fileName)
	if err != nil {
		return fmt.Sprintf("Error opening file %s: %v", fileName, err)
	}

	// check if the file exist
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		return fmt.Sprintf("Can't find '%s'.", fileName)
	}

	// check for duplicates
	output := fmt.Sprintf("'%s'is already tracked.", fileName)
	if containsFile(u.FileNames, fileName) {
		return output
	}

	// add to slice of tracked files
	u.Files = append(u.Files, openedFile)
	// add to tracked file names
	u.FileNames = append(u.FileNames, fileName)
	return fmt.Sprintf("The file '%v' is tracked.", fileName)
}
