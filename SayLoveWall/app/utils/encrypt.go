package utils

import (
	"crypto/sha256"
	"encoding/hex"
)


func Encrypt(str string) string { //Sha256加密方式
	srcByte := []byte(str)
	sha256New := sha256.New()
	sha256Bytes := sha256New.Sum(srcByte)
	sha256String := hex.EncodeToString(sha256Bytes)
	return sha256String
}