package goutil

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func Base64Encode(str string) string {
	s := []byte(str)
	return base64.StdEncoding.EncodeToString(s)
}

func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
