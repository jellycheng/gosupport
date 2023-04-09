package xcrypto

import (
	"encoding/base64"
	"strings"
)

func Base64StdEncode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64StdDecode(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

// Base64UrlSafeDecode Base64安全解密
func Base64UrlSafeDecode(s string) (string, error) {
	var c = (4 - len(s)%4) % 4
	s += strings.Repeat("=", c)
	res, err := base64.URLEncoding.DecodeString(s)
	return string(res), err
}

// Base64UrlSafeEncode Base64安全加密
func Base64UrlSafeEncode(s string) string {
	bytearr := base64.StdEncoding.EncodeToString([]byte(s))
	ret := strings.Replace(string(bytearr), "/", "_", -1)
	ret = strings.Replace(ret, "+", "-", -1)
	ret = strings.Replace(ret, "=", "", -1)
	return ret
}
