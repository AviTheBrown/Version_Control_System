package datatypes

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var err error
	// pwd
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}
	// return to parent dir
	defer func() {
		os.Chdir(currentDir)
		if err != nil {
			fmt.Println("Error returning to original dir:", err)
		}
	}()

	// move down to ./vcs
	targetDir := "./vcs"
	err = os.Chdir(targetDir)
	if err != nil {
		log.Fatal(err)
	}
	// check if the file is in ./vcs
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
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

// func containsFile(searchedFile, indexedFile string) string {
// 	file, err := os.Open(searchedFile)
// 	if err != nil {
// 		return indexedFile
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if strings.Contains(line, indexedFile) {

// 		}
// 	}
// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("there was a problem scanning the file.")
// 	}
// 	return true
// }

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
func (u *User) AddAction(fileName string) {

	err := os.Chdir("./vcs")
	if err != nil {
		fmt.Println("there was a problem here")
	}
	defer os.Chdir("..")

	writeToFile := func() {
		err := os.WriteFile("index.txt", []byte(fileName+"\n"), 0644)
		if err != nil {
			fmt.Printf("unable to write to file")
		}
	}
	writeToFile()

	// for looking up a file name in index.tx
	searchForFile := func() bool {
		_, err := os.Open(fileName)
		if err != nil {
			formatOutput(fileName, false)
			return false
		}
		formatOutput(fileName, true)
		return true
	}
	searchForFile()
	displayFiles(fileName)
}
