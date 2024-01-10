package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type SVCS map[string]string

func printValidCommands(wrongCmd string, mySVCS SVCS) string {
	var commandBuilder strings.Builder
	commandBuilder.WriteString(fmt.Sprintf("'%s' is not a valid SVCS command.", wrongCmd))

	// for commands, description := range mySVCS {
	// 	commandBuilder.WriteString(fmt.Sprintf("%-10s%s\n", commands, description))
	// }
	return commandBuilder.String()
}

func commandDescription(command string, svc SVCS) string {
	if description, ok := svc[command]; ok {
		return description
	} else {
		return printValidCommands(command, svc)
	}
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
	flag.StringVar(&command, "command", "", printValidCommands(command, mySVCS))
	flag.Parse()
	description := commandDescription(command, mySVCS)
	fmt.Println(description)
}
