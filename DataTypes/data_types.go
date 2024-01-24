package datatypes

import (
	mutex "Version_Control_System/Mutex"
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
func (u *User) AddAction(fileName string) (updatedUser *User, output string) {
	mutex := mutex.GetUserMutexData()
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println("before")
	if u.isFileTracked(fileName) {
		fmt.Println("After(tracked)")
		output = formatOutput(fileName, true)
		return u, output
	}
	fmt.Println("After (untracked)")
	u.FileNames = append(u.FileNames, fileName)
	output = formatOutput(fileName, false)
	updatedUser = u
	return
}
