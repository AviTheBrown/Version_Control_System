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

func (u *User) LoadTrackedFiles(filename string) []string {
	content, err := os.ReadFile("vcs/index.txt")
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
	appendToFile("vcs/index.txt", filename)
	u.Files = append(u.Files, filename)
	fmt.Println(formatOutput(filename, false))

}
func (u *User) isFileTracked(filename string) bool {
	trackedFiles := u.LoadTrackedFiles("vcs/index.txt")
	for _, trackedFile := range trackedFiles {
		if trackedFile == filename {
			return true
		}
	}
	return false
}

func appendToFile(filename, content string) {
	file, err := os.OpenFile("./vcs/index.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error writing to file %v\n", err)
	}
	defer file.Close()
	_, err = fmt.Fprint(file, content, "\n")
	if err != nil {
		log.Fatal(err)
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
