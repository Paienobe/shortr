package utils

import (
	"crypto/sha256"
	"fmt"
)

func GenerateHash(url string) string {
	hash := sha256.New()
	hash.Write([]byte(url))
	final := hash.Sum(nil)[:3]
	return fmt.Sprintf("%x", final)
}
