package datatypes

import (
	"fmt"
)

type SVCS map[string]string

type FileInfo struct {
	FileNames []string
}
type User struct {
	UserName string
	FileInfo
}

func CreateUser() *User {
	return &User{
		UserName: "Default Name",
		FileInfo: FileInfo{
			FileNames: make([]string, 0),
		},
	}
}
func (u User) isFileTracked(filename string) (isTracked bool) {
	for _, existingFile := range u.FileNames {
		if filename == existingFile {
			return
		}
	}
	return
}

func formatOutput(fileName string, isTracked bool) string {
	if isTracked {
		return fmt.Sprintf("'%s' is alread tracked.", fileName)
	}
	return fmt.Sprintf("The file '%s' is now tracked.", fileName)
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
func (u *User) AddAction(fileName string) (updatedUser *User, output string) {
	if u.isFileTracked(fileName) {
		return u, formatOutput(fileName, true)
	}
	u.FileNames = append(u.FileNames, fileName)
	updatedUser = u
	// fmt.Println("tesssstttt")
	return updatedUser, formatOutput(fileName, false)
}
