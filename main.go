package main

import (
	"Version_Control_System/DataTypes"
	"flag"
	"fmt"
	"os"
)

var user *datatypes.User
var commandOrder []string

func main() {
	mySVCS := datatypes.SVCS{
		"config":   "Get and set a username.",
		"add":      "Add a file to the index.",
		"log":      "Show commit logs.",
		"commit":   "Save changes.",
		"checkout": "Restore a file.",
	}

	commandOrder = []string{"config", "add", "log", "commit", "checkout"}

	user = datatypes.LoadUser()
	processCommandLine(mySVCS, commandOrder)
}

func commandActions(command string, usr *datatypes.User) {
	switch command {
	case "config":
		fmt.Println(usr.ConfigAction(os.Args[1]))
	case "add":
		if len(os.Args) > 2 {
			result := usr.AddAction(os.Args[2])
			fmt.Println(result)
			datatypes.SaveUser(user)
		}
	default:
		fmt.Println("TODO")
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

	commandActions(command, user)

	if command == "add" {
		// if its the first time using add ./main add
		if len(user.FileNames) == 0 {
			printValidCommands(command, mySCVS)
			// if only the add command is used with tracked filed ./main add
		} else if len(user.FileNames) > 0 && len(flag.Args()) == 1 {
			// prints out all the tracked files
			fmt.Println("Tracked files:")
			for _, file := range user.FileNames {
				fmt.Println(file)
			}
		} else {

		}
	}
	// // if command == "add" {
	// // 	if len(user.FileNames) == 0 {
	// 		// printValidCommands(command, mySCVS)
	// } else if len(user.FileNames) > 0 && flag.Arg(1) == ""{
	// 			print
	// }
	// 	} else if len(user.FileNames) > 1 {
	// 		fmt.Println("Tracked files:")
	// 		for _, file := range user.FileNames {
	// 			fmt.Println(file)
	// 		}
	// 	}
	// } else {
	// 	printValidCommands(command, mySCVS)
	// }
	// printValidCommands(command, mySCVS)
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
