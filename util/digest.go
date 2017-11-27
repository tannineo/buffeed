package util

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
)

var md5Digester hash.Hash

func init() {
	md5Digester = md5.New()
}

// GetMd5String 生成32位md5字串
func GetMd5String(s string) string {
	md5Digester.Reset()
	md5Digester.Write([]byte(s))
	return hex.EncodeToString(md5Digester.Sum(nil))
}
