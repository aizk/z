package zcrypto

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(args ...string) string {
	h := md5.New()
	h.Write([]byte(strings.Join(args, "")))
	return hex.EncodeToString(h.Sum(nil))
}
