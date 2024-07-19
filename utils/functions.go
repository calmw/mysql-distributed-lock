package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"strconv"
	"time"
)

// GetMd5String 生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// UniqueId 生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	encryptedString := GetMd5String(base64.URLEncoding.EncodeToString(b))
	return encryptedString[0:16] + Int64ToString(time.Now().Unix()) + encryptedString[26:]
}

// Int64ToString int64转字符串
func Int64ToString(n int64) string {
	i := int64(n)
	return strconv.FormatInt(i, 10)
}
