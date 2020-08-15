package curl

import (
	"encoding/base64"
	"net/url"
)

func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

func UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) (string, error) {
	s, err := base64.StdEncoding.DecodeString(str)
	return string(s), err
}

