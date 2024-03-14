package hashing

import (
	datatypes "Version_Control_System/DataTypes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func GenerateHashForFiles(files []string, user datatypes.User) string {
	sha256Hash := sha256.New()
	var fileHashString string
	for _, file := range files {
		file, err := os.Open(file)
		if err != nil {
			fmt.Println("There was a problem opening the file.")
			log.Fatal(err)
		}
		if _, err := io.Copy(sha256Hash, file); err != nil {
			fmt.Println("There was a problem hashing the file.")
			log.Fatal()
		}
		sha256HashValue := sha256Hash.Sum(nil)
		fileHashString = hex.EncodeToString(sha256HashValue)
		user.FileInfo.FileHash = fileHashString

	}
	return fileHashString
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
