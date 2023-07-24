package xcrypto

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5To16(s string) string {
	ret := MD5(s)
	return ret[8:24]
}
