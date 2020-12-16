package curl

import (
	"encoding/base64"
	"net/url"
	"strings"
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

func TrimPath(path string, pos int) string {
	ret := ""
	switch pos {
	case 0:
		ret = strings.Trim(path, "/")
	case 1:
		ret = strings.TrimLeft(path, "/")
	case 2:
		ret = strings.TrimRight(path, "/")
	default:
		ret = path
	}
	return ret
}
