package datatypes

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
func isFileTracked(filename string) bool {
	// Move down to ./vcs
	// Read the content of index.txt
	fileContent, err := os.ReadFile("index.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Check if the filename is present in the content
	return strings.Contains(string(fileContent), filename)
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
func (u *User) AddAction(fileName string) {
	err := os.Chdir("./vcs")
	if err != nil {
		fmt.Println("there was a problem here")
	}
	defer os.Chdir("..")

	if isFileTracked(fileName) {
		return
	} else {
		u.FileNames = append(u.FileNames, fileName)
	}
	writeToFile := func() {
		file, err := os.OpenFile("index.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Printf("unable to write to file")
		}
		defer file.Close()
		data := []byte(fileName)
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	writeToFile()
}
