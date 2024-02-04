package datatypes

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
		Files:    loadTrackedFiles("vcs/index.txt"),
	}
	return user
}

func loadTrackedFiles(filename string) []string {
	content, err := os.ReadFile("vcs/index.txt")
	if err != nil {
		fmt.Printf("Error reading tracked files %v\n", err)
		return nil
	}
	return strings.Split(string(content), "\n")
}
func (u *User) AddAction(filename string) {
	if isFileTracked(filename) {
		fmt.Println(formatOutput(filename, true))
		return
	}
	u.Files = append(u.Files, filename)
	appendFile("vcs/index.txt", filename)
	fmt.Println(formatOutput(filename, false))

}
func isFileTracked(filename string) bool {
	for _, trackedFile := range loadTrackedFiles("vcs/index.txt") {
		// if the base files are the same then the file is tracked.
		if filepath.Base(trackedFile) == filepath.Base(filename) {
			return true
		}
	}
	return false
}
func appendFile(filename, content string) {
	err := os.WriteFile(filename, []byte(content), os.ModeAppend)
	if err != nil {
		fmt.Printf("Error writing to file %v\n", err)
	}
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

func DisplayFiles(filename string) {
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
