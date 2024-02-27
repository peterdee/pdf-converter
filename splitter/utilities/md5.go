package utilities

import (
	"crypto/md5"
	"encoding/hex"
)

func CreateMD5Hash(plaintext string) string {
	if plaintext == "" {
		return ""
	}
	hash := md5.Sum([]byte(plaintext))
	return hex.EncodeToString(hash[:])
}
