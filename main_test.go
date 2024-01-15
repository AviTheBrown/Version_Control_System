package main

import (
	datatypes "Version_Control_System/DataTypes"
	"os"
	"testing"
)

func TestPrintedValues(t *testing.T) {
	mySVCS := datatypes.SVCS{
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

func TestCreateUser(t *testing.T) {
	user1 := new(datatypes.User)
	expectedUserDefaultName := ""
	actualUserDefaultName := user1.UserName

	if expectedUserDefaultName != actualUserDefaultName {
		t.Errorf("Expected :%q but got :%q", expectedUserDefaultName, actualUserDefaultName)
	}

	expectedUserDefaultFile := (*os.File)(nil)
	actualUserDefaultFile := user1.FileInfo.File

	if expectedUserDefaultFile != actualUserDefaultFile {
		t.Errorf("Expected File to be : %v but got %v", expectedUserDefaultFile, actualUserDefaultFile)
	}

}
