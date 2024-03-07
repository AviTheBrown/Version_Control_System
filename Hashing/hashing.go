package hashing

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
)

func GenerateCommitHashID(commitMessage string) (string, error) {
	if commitMessage == "" {
		return "", errors.New("Commit messages cannot be empty.")
	}
	timeStamp := time.Now().Format(time.RFC3339)
	hashInput := commitMessage + timeStamp

	hash := sha256.Sum256([]byte(hashInput))
	commitID := fmt.Sprint("%x", hash)

	return commitID, nil
}
func HashComparison(hash1, hash2 string) bool {
	return hash1 == hash2
}
