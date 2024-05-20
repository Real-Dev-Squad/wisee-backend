package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateHash(content string, length int8) (string, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(content))

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil))[:length], nil
}
