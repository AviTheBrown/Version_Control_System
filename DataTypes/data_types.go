package datatypes

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type SVCS map[string]string

type User struct {
	UserName string
	Files    []string
}

func CreateUser() *User {
	user := &User{
		UserName: "Name",
	}
	return user
}

const (
	configFilePath = "vcs/config.txt"
	indexFilePath  = "vcs/index.txt"
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
	return strings.Split(string(content), "\n")
}
func (u *User) AddAction(filename string) {
	if u.isFileTracked(filename) {
		fmt.Println(formatOutput(filename, true))
		return
	}
	appendToFile(indexFilePath, filename)
	u.Files = append(u.Files, filename)
	fmt.Println(formatOutput(filename, false))

}
func (u *User) isFileTracked(filename string) bool {
	trackedFiles := u.LoadTrackedFiles(indexFilePath)
	for _, trackedFile := range trackedFiles {
		if trackedFile == filename {
			return true
		}
	}
	return false
}

func appendToFile(filePath, content string) {
	switch {
	case filePath == configFilePath:
		err := os.WriteFile(filePath, []byte(content), os.ModePerm)
		if err != nil {
			fmt.Println("Error writing to file.")
		}
	case filePath == indexFilePath:

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
	appendToFile(configFilePath, userName)
	return fmt.Sprintf("The username is %s.", userName)
}

func displayFiles(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	fmt.Println("Tracked files:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func creatFile(filename string) {
	_, err := os.Create(filename)
	if err != nil {
		log.Fatal("there was a error trying to create the file here:", err)
	}
}
