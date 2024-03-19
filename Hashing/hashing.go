package hashing

import (
	datatypes "Version_Control_System/DataTypes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

func GenerateCommitHashID(commitMessage string) (string, error) {
	if commitMessage == "" {
		return "Message was not passed.", nil
	}
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(commitMessage))

	hexString := hex.EncodeToString(sha256Hash.Sum(nil))
	return hexString, nil
}
func getLastCommitHash(user datatypes.User) (string, error) {
	if len(user.FileMeta) >= 1 {
		fmt.Println(len(user.FileMeta))
		lastFileHash := user.FileMeta[len(user.FileMeta)-1].FileHash
		fmt.Printf("The last file hased is: %s", lastFileHash)
		return lastFileHash, nil
	} else {
		return "there are no commits yet.", fmt.Errorf("There are no files avaliable.")
	}
}
func CheckForChanges(user datatypes.User) bool {
	for _, file := range user.FileMeta {
		trackedFilePath := filepath.Join("./", file.FileName)
		fileData, err := os.ReadFile(trackedFilePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", trackedFilePath, err)
			continue
		}
		currentHash := HashFileData(fileData)
		if file.FileHash != currentHash {
			return true // File has changed
		}
	}

	return false // No changes detected

}
func HashComparison(hash1, hash2 string) bool {
	return hash1 == hash2
}

func HashFileData(data []byte) string {
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	hash := sha256Hash.Sum(nil)
	return hex.EncodeToString(hash)
}
