package xcrypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
)

func HmacMd5(dataStr, key string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(dataStr))
	return hex.EncodeToString(h.Sum([]byte("")))
}

func HmacSha1(dataStr, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(dataStr))
	return hex.EncodeToString(h.Sum([]byte("")))
}

func HmacSha512(dataStr, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(dataStr))
	return hex.EncodeToString(h.Sum([]byte("")))
}
