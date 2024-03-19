package main

import (
	datatypes "Version_Control_System/DataTypes"
	files "Version_Control_System/Files"
	hashing "Version_Control_System/Hashing"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	// "strings"
)

var user *datatypes.User
var commandOrder = []string{"config", "add", "log", "commit", "checkout"}

func main() {
	user, _ = datatypes.CreateUser()

	if !files.AllFilesCreated() {
		files.CreateVCSDirWithChildFiles()
	}
	trackedFilesInIndexFile := user.LoadTrackedFiles("vcs/index.txt")

	for _, file := range trackedFilesInIndexFile {
		fileData, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %v:", file)
			continue
		}
		sha256Hash := hashing.HashFileData(fileData)
		user.AddFileToMeta(file, fileData, sha256Hash)
	}
	mySVCS := datatypes.SVCS{
		"config":   "Get and set a username.",
		"add":      "Add a file to the index.",
		"log":      "Show commit logs.",
		"commit":   "Save changes.",
		"checkout": "Restore a file.",
	}
	processCommandLine(mySVCS, commandOrder)
}

func commandActions(command string, usr *datatypes.User, mySVCS datatypes.SVCS) {
	switch command {
	case "config":
		content, _ := os.ReadFile("vcs/config.txt")
		if len(os.Args) > 2 {
			fmt.Println(user.ConfigAction(os.Args[2]))
		} else if len(content) != 0 {
			fmt.Printf("The username is %s.", content)
		}
	case "add":
		if len(os.Args) > 2 {
			var result string
			user.AddAction(os.Args[2])
			fmt.Println(result)
		} else {
		}
	case "commit":
		var (
			commitMessage string
			hashString    string
			err           error
		)
		if len(os.Args) > 2 {
			commitMessage = strings.Join(os.Args[2:], " ")
			hashString, err = hashing.GenerateCommitHashID(commitMessage)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Message not sent.")
			return
		}
		fmt.Println("test reach")
		hasFileChanged := hashing.CheckForChanges(*user)
		if hasFileChanged {
			files.CreateHashDir(commitMessage, hashString, *user)
			files.LogAction(hashString, user.UserName, commitMessage)
		} else {
			fmt.Println("Nothing to commit (main)")
		}

	case "log":
		fileData, err := os.Stat(datatypes.LOGFILEPATH)
		if err != nil {
			fmt.Println("Error finding file.")
			return
		}
		if fileData.Size() == 0 {
			fmt.Println("No commits yet.")
		} else {
			contents, err := os.ReadFile(datatypes.LOGFILEPATH)
			if err != nil {
				fmt.Println("there was an error reading the file.")
			}
			fmt.Println(string(contents))
		}

	default:
		defaultString := fmt.Sprintf(printValidCommands(command, mySVCS))
		fmt.Println(defaultString)
	}
}
func processCommandLine(mySCVS datatypes.SVCS, svcsOrder []string) {
	helpflag := flag.Bool("help", false, "Prints Help Message")
	flag.Parse()

	if *helpflag || flag.NArg() == 0 || flag.Arg(0) == "help" {
		printAllCommands(mySCVS, svcsOrder)
		return
	}
	command := flag.Arg(0)
	commandActions(command, user, mySCVS)
	if command == "add" {
		switch {
		// if there are tracked files but you only wish to display them
		// with only the "add" command i.e ./main add
		case len(user.FileMeta) == 0 && flag.NArg() == 1:
			fmt.Println(printValidCommands(command, mySCVS))
			// if there are file that are added and you wish to display them to stout
		case len(user.FileMeta) > 0 && flag.NArg() == 1:
			fmt.Println("Tracked Files:")
			for _, file := range user.FileMeta {
				fileData := file.FileName
				fmt.Println(fileData)
			}
		}
	}
}

func printAllCommands(mySVCS datatypes.SVCS, svcsOrder []string) {
	fmt.Println("These are SVCS commands:")
	for _, command := range svcsOrder {
		decription := mySVCS[command]
		fmt.Printf("%-10s%s\n", command, decription)
	}
}

func printValidCommands(command string, mySVCS datatypes.SVCS) string {
	if _, ok := mySVCS[command]; ok {
		return commandDescription(command, mySVCS)
	} else {
		return fmt.Sprintf("'%s' is not a SVCS command.", command)
	}
}

func commandDescription(command string, svc datatypes.SVCS) string {
	if description, ok := svc[command]; ok {
		return description
	} else {
		return fmt.Sprintf("'%s' is not a SVCS command.", command)
	}
}
