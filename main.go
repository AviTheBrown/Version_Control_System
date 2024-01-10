package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type SVCS map[string]string

func main() {

	mySVCS := SVCS{
		"config":   "Get and set a username.",
		"add":      "Add a file to the index.",
		"log":      "Show commit logs.",
		"commit":   "Save changes.",
		"checkout": "Restore a file."}

	command := flag.String("command", "", "These are SVCS commands:")
	flag.Parse()

	if *command == "help" || *command == "" {
		printValidCommands(mySVCS)
		os.Exit(0)
	}

	description, err := validateCommand(*command, mySVCS)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("%s %s", *command, description)
}
func validateCommand(command string, mySVCS SVCS) (string, error) {
	if command == "help" || command == "" {
		return "", fmt.Errorf(printValidCommands(mySVCS))
	}
	value, ok := mySVCS[command]
	if !ok {
		printValidCommands(mySVCS)
	}
	return value, nil
}

func printValidCommands(mySVCS SVCS) string {
	var commandBuilder strings.Builder
	commandBuilder.WriteString("These are SVCS commands:\n")

	for commands, description := range mySVCS {
		commandBuilder.WriteString(fmt.Sprintf("%s\t%s\n", commands, description))
	}
	return commandBuilder.String()
}
