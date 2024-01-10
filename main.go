package main

import (
	"flag"
	"fmt"
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
	result, err := validateCommand(*command, mySVCS)
	if err != nil {
		fmt.Println(err)
		optionString := printValidCommands(mySVCS)
		fmt.Println(optionString)
		return
	}

	fmt.Println(result)
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
	commandBuilder.WriteString("These are SVCS commands:")

	for commands, description := range mySVCS {
		commandBuilder.WriteString(fmt.Sprintf("%s\t%s\n", commands, description))
	}
	return commandBuilder.String()
}
