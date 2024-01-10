package main

import (
	"fmt"
	"os"
)

type SVCS map[string]string

func printAllCommands(mySVCS SVCS, svcsOrder []string) {
	fmt.Println("These are SVCS commands:")
	for _, command := range svcsOrder {
		decription := mySVCS[command]
		fmt.Printf("%-10s%s\n", command, decription)
	}
}
func printValidCommands(command string, mySVCS SVCS) string {
	if _, ok := mySVCS[command]; ok {
		return commandDescription(command, mySVCS)
	} else {
		return fmt.Sprintf("'%s' is not a SVCS command.", command)
	}
}

func commandDescription(command string, svc SVCS) string {
	if description, ok := svc[command]; ok {
		return description
	} else {
		return fmt.Sprintf("'%s' is not a SVCS command.", command)
	}
}

func main() {
	mySVCS := SVCS{
		"config":   "Get and set a username.",
		"add":      "Add a file to the index.",
		"log":      "Show commit logs.",
		"commit":   "Save changes.",
		"checkout": "Restore a file.",
	}

	commandOrder := []string{"config", "add", "log", "commit", "checkout"}

	// Check if there are enough command-line arguments
	if len(os.Args) < 2 || os.Args[1] == "help" || os.Args[1] == "-help" || os.Args[1] == "--help" {
		printAllCommands(mySVCS, commandOrder)
		return
	}
	// Retrieve the command from command-line arguments
	command := os.Args[1]

	// Use the command directly without using flag.StringVar
	description := printValidCommands(command, mySVCS)

	// Print the description
	fmt.Println(description)
}
