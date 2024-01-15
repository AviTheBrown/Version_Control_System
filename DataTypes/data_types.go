package datatypes

import (
	"fmt"
	"os"
)

type SVCS map[string]string

type FileInfo struct {
	File     *os.File
	FileName string
}
type User struct {
	UserName string
	FileInfo
}

func (u *User) ConfigAction(userName string) *User {
	if u.UserName == userName {
		//  create a temp User object with only UserName field
		return &User{UserName: userName}
	}
	u.UserName = userName
	return u
}

func (u *User) AddAction() {
	var fileName string
	fmt.Scan(&fileName)
	u.FileInfo.FileName = fileName
}
