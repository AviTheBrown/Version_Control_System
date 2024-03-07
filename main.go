package main

import (
	datatypes "Version_Control_System/DataTypes"
	files "Version_Control_System/Files"
	"flag"
	"fmt"
	"os"
)

var user *datatypes.User
var commandOrder = []string{"config", "add", "log", "commit", "checkout"}

func main() {
	user = datatypes.CreateUser()

	if _, err := os.Stat("vcs"); os.IsNotExist(err) {
		files.CreateDirWithChildFiles()
	}
	user.FileInfo.FileNames = user.LoadTrackedFiles("vcs/index.txt")
	user.LoadUserName("vcs/config.txt")

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
		fmt.Println("test")
		// if len(os.Args) > 2 {
		// 	commitMessage := os.Args[2]
		// 	hashing.GenerateCommitHashID(commitMessage)

		// }

	case "log":
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
		case len(user.FileInfo.FileNames) == 0 && flag.NArg() == 1:
			fmt.Println(printValidCommands(command, mySCVS))
			// if there are file that are added and you wish to display them to stout
		case len(user.FileInfo.FileNames) > 0 && flag.NArg() == 1:
			fmt.Println("Tracked Files:")
			for _, file := range user.FileInfo.FileNames {
				fmt.Println(file)
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
