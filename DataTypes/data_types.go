package datatypes

import (
	"fmt"
)

type SVCS map[string]string

type FileInfo struct {
	// Files     []*os.File
	FileNames []string
}
type User struct {
	UserName string
	FileInfo
}

func CreateUser() *User {
	return &User{}
}
func (u User) isFileTracked(filename string) bool {
	for _, existingFile := range u.FileNames {
		if filename == existingFile {
			return true
		}
	}
	return false
}

func formatOutput(fileName string, isTracked bool) string {
	if isTracked {
		return fmt.Sprintf("'%s' is alread tracked.", fileName)
	}
	return fmt.Sprintf("The file '%s' is tracked.", fileName)
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
func (u User) AddAction(fileName string) (updatedUser User, output string) {

	// // create the file
	// openedFile, err := files.CreateFile(fileName)
	// if err != nil {
	// 	return fmt.Sprintf("Error opening file %s: %v", fileName, err)
	// }

	// check if the file exist
	// _, err = os.Stat(fileName)
	// if os.IsNotExist(err) {
	// 	return fmt.Sprintf("Can't find '%s'.", fileName)
	// }

	// check for duplicates
	if u.isFileTracked(fileName) {
		return u, formatOutput(fileName, true)
	}
	updatedUser = User{
		UserName: u.UserName,
		FileInfo: FileInfo{FileNames: append(u.FileNames, fileName)},
	}
	output = formatOutput(fileName, false)
	return
}
