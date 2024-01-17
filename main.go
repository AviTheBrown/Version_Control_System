package main

import (
	"Version_Control_System/DataTypes"
	"flag"
	"fmt"
	"os"
)

func main() {
	mySVCS := datatypes.SVCS{
		"config":   "Get and set a username.",
		"add":      "Add a file to the index.",
		"log":      "Show commit logs.",
		"commit":   "Save changes.",
		"checkout": "Restore a file.",
	}

	commandOrder := []string{"config", "add", "log", "commit", "checkout"}

	processCommandLine(mySVCS, commandOrder)
}

func commandActions(command string, usr datatypes.User) {
	switch command {
	case "config":
		fmt.Println(usr.ConfigAction(os.Args[2]))
	case "add":
		if len(os.Args) > 2 {
			fmt.Println(usr.AddAction(os.Args[2]))
		} else {
			fmt.Println("Add a file to the index.")
		}

	}
}
func processCommandLine(mySCVS datatypes.SVCS, svcsOrder []string) {
	user := datatypes.CreateUser()
	helpflag := flag.Bool("help", false, "Prints Help Message")
	flag.Parse()

	if *helpflag || flag.NArg() == 0 || flag.Arg(0) == "help" {
		printAllCommands(mySCVS, svcsOrder)
		return
	}
	command := flag.Arg(0)

	commandActions(command, *user)

	if command == "add" && len(user.FileNames) > 0 {
		fmt.Println("Tracked Files:")
		for _, file := range user.FileNames {
			fmt.Println(file)
		}
	} else {
		printValidCommands(command, mySCVS)
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
