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

func commandActions(command string, usr datatypes.User) *datatypes.User {
	switch command {
	case "config":
		usr.ConfigAction(command)
	case "add":
		usr.AddAction(command)

	}
	return &usr
}

// func flagProcessing(flag *string) {
// 	switch flag{
// 		case "add":

//		}
//	}
func processCommandLine(mySCVS datatypes.SVCS, svcsOrder []string) {
	helpflag := flag.Bool("help", false, "Prints Help Message")
	flag.Parse()

	if *helpflag || flag.NArg() == 0 || flag.Arg(0) == "help" {
		printAllCommands(mySCVS, svcsOrder)
		return
	}

	command := flag.Arg(0)
	description := printValidCommands(command, mySCVS)

	user := datatypes.CreateUser()
	if len(os.Args) >= 3 {
		switch os.Args[1] {
		case "add":
			filepath := os.Args[2]
			fmt.Println(user.AddAction(filepath))
		}
	}
	fmt.Println(description)
	fmt.Println(user.FileInfo)
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
