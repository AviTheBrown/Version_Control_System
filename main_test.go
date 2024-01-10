package main

import "testing"

func TestPrintedValues(t *testing.T) {
	mySVCS := SVCS{
		"config":   "Get and set a username.",
		"add":      "Add a file to the index.",
		"log":      "Show commit logs.",
		"commit":   "Save changes.",
		"checkout": "Restore a file.",
	}
	valideCommand := "config"
	expectedResult := "Get and set a username."
	actualResult := printValidCommands(valideCommand, mySVCS)

	if expectedResult != actualResult {
		t.Errorf("Expected output %q but got %q", expectedResult, actualResult)
	}

	invalidCommand := "invalid"
	expectedError := "'invalid' is not a SVCS command."
	actualError := printValidCommands(invalidCommand, mySVCS)

	if expectedError != actualError {
		t.Errorf("Expected error: %q but got %q", expectedError, actualError)
	}
}
