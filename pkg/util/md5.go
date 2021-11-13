package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 格式化上传文件     防止将原始名称暴露
func EncodeMD5(value string)string{
	m:=md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}