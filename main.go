package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type SVCS map[string]string

func printValidCommands(mySVCS SVCS) string {
	var commandBuilder strings.Builder
	commandBuilder.WriteString("These are SVCS commands:\n")

	for commands, description := range mySVCS {
		commandBuilder.WriteString(fmt.Sprintf("%-10s%s\n", commands, description))
	}
	return commandBuilder.String()
}

func commandDescription(command string, svc SVCS) string {
	if description, ok := svc[command]; ok {
		fmt.Println(description)
		return description
	}
	return printValidCommands(svc)
}
func parseCommand(command string, svc SVCS) string {
	if description, ok := svc[command]; ok {
		return description
	}
	return "help"
}
func main() {

	mySVCS := SVCS{
		"config":   "Get and set a username.",
		"add":      "Add a file to the index.",
		"log":      "Show commit logs.",
		"commit":   "Save changes.",
		"checkout": "Restore a file."}

	// reterive the [1] arg to pass as the flag
	command := os.Args[1]
	// trimmedCommad := command[1:]
	// cmd := flag.String("command", "", printValidCommands(mySVCS))
	flag.StringVar(&command, "command", "", printValidCommands(mySVCS))
	flag.Parse()
	commandDescription(command, mySVCS)
}
