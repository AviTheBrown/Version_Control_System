package main

import (
	datatypes "Version_Control_System/DataTypes"
	mutex "Version_Control_System/Mutex"
	"flag"
	"fmt"
	"os"
)

var user *datatypes.User
var commandOrder = []string{"config", "add", "log", "commit", "checkout"}

func main() {
	user = datatypes.CreateUser()
	fmt.Println(user)
	mySVCS := datatypes.SVCS{
		"config":   "Get and set a username.",
		"add":      "Add a file to the index.",
		"log":      "Show commit logs.",
		"commit":   "Save changes.",
		"checkout": "Restore a file.",
	}

	mutex.GetUserMutexData()
	processCommandLine(mySVCS, commandOrder)
}

func saveUserData() {
	mutex.UserDataMutex.Lock()
	defer mutex.UserDataMutex.Unlock()
}
func commandActions(command string, usr *datatypes.User, mySVCS datatypes.SVCS) {
	switch command {
	case "config":
		fmt.Println(usr.ConfigAction(os.Args[1]))
	case "add":
		if len(os.Args) > 2 {
			var result string
			mutex.UserDataMutex.Lock()
			defer mutex.UserDataMutex.Unlock()
			updatedUser, result := user.AddAction(os.Args[2])
			user = updatedUser
			fmt.Println(result)
			fmt.Println(user)
		} else {
			printValidCommands(command, mySVCS)
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

	saveUserData()
	mutex.UserDataMutex.Lock()
	defer mutex.UserDataMutex.Unlock()
	if command == "add" {
		// if its the first time using add ./main add

		if len(user.FileNames) > 0 {
			fmt.Println("Tracked Files:")
			for _, file := range user.FileNames {
				fmt.Println(file)
			}
		} else {
			fmt.Println("No tracked files")

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
