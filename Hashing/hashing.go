package hashing

import (
	"crypto/sha256"
	"encoding/hex"
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
func HashComparison(hash1, hash2 string) bool {
	return hash1 == hash2
}
