package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5String(s string) string {
	if s == "" {
		return ""
	}

	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
