package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 对文件名用 MD5 加密后再进行写入 避免直接暴露原始名称
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
