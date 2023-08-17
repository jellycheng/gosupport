package xcrypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
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

func HmacSHA256(data, key string) string {
	var h = hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func HmacSHA256ForByte(data, key []byte) []byte {
	var h = hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func HmacSha512(dataStr, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(dataStr))
	return hex.EncodeToString(h.Sum([]byte("")))
}
