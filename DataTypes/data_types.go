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
	return new(User)
}
func (u *User) ConfigAction(userName string) string {
	if u.UserName == userName {
		//  create a temp User object with only UserName field
		return fmt.Sprintf("The username is %s", u.UserName)
	}
	u.UserName = userName
	return fmt.Sprintf("The username is %s", u.UserName)
}

func (u *User) AddAction(fileName string) string {
	u.FileNames = append(u.FileNames, fileName)
	return fmt.Sprintf("The file '%v' is tracked.", fileName)
}
