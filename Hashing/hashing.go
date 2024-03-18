package hashing

import (
	datatypes "Version_Control_System/DataTypes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func gernerateHashForFileMinor(fileData []byte) (string, error) {
	sha256Hash := sha256.New()
	sha256Hash.Write(fileData)
	sha256HashValue := sha256Hash.Sum(nil)
	filehash := hex.EncodeToString(sha256HashValue)
	return filehash, nil
}
func GenerateHashForFiles(files []string, user *datatypes.User) {
	for _, file := range files {
		fileData, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("There was a problem opening the file.")
			log.Fatal(err)
		}

		fileHashString, err := gernerateHashForFileMinor(fileData)
		if err != nil {
			fmt.Println("There was a problem create the hash.")
			log.Fatal(err)
		}
		user.AddFileToMeta(file, fileData, fileHashString)
	}
}
func GenerateCommitHashID(commitMessage string) (string, error) {
	if commitMessage == "" {
		return "Message was not passed.", nil
	}
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(commitMessage))

	hexString := hex.EncodeToString(sha256Hash.Sum(nil))
	return hexString, nil
}
func HashComparison(hash1, hash2 string) bool {
	return hash1 == hash2
}
