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
	return &User{}
}
func (u *User) ConfigAction(userName string) string {
	if u.UserName == userName {
		//  create a temp User object with only UserName field
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
	output := fmt.Sprintf("'%s'is already tracked.", fileName)
	if containsFile(u.FileNames, fileName) {
		for _, files := range u.FileNames {
			fmt.Println(files)
		}
		return output
	}
	u.FileNames = append(u.FileNames, fileName)
	return fmt.Sprintf("The file '%v' is tracked.", fileName)
}
