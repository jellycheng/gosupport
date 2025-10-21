package curl

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"unicode"
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

func Base64EncodeV2(b []byte) []byte {
	buf := make([]byte, base64.RawURLEncoding.EncodedLen(len(b)))
	base64.RawURLEncoding.Encode(buf, b)
	return buf
}

func Base64DecodeV2(b []byte) ([]byte, error) {
	buf := make([]byte, base64.RawURLEncoding.DecodedLen(len(b)))
	n, err := base64.RawURLEncoding.Decode(buf, b)
	return buf[:n], err
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

// 向url中追加参数
func AppendParamToUrl(url, param string) string {
	if IsBlank(param) {
		return url
	}
	tmpParam := []rune(param)

	questionMarkIndex := strings.Index(url, "?")
	alarmMarkIndex := strings.Index(url, "#")
	if questionMarkIndex == -1 && alarmMarkIndex == -1 { // 不存在#、?
		if tmpParam[0] == '#' || tmpParam[0] == '?' {
			return fmt.Sprintf("%s%s", url, param)
		}
		return fmt.Sprintf("%s?%s", url, param)
	}
	if questionMarkIndex >= 0 && alarmMarkIndex == -1 { // 仅存在?
		if tmpParam[0] == '#' {
			return fmt.Sprintf("%s%s", url, param)
		} else if tmpParam[0] == '?' {
			return fmt.Sprintf("%s&%s", url, strings.TrimLeft(param, "?"))
		}
		return fmt.Sprintf("%s&%s", url, param)
	}
	if questionMarkIndex == -1 && alarmMarkIndex >= 0 { // 仅存在#
		tmpUrl := strings.SplitN(url, "#", 2)
		return fmt.Sprintf("%s?%s#%s", tmpUrl[0], strings.TrimLeft(param, "?#"), tmpUrl[1])
	}
	// 存在?#
	if questionMarkIndex < alarmMarkIndex { //?在#号之前
		tmpUrl := strings.SplitN(url, "#", 2)
		return fmt.Sprintf("%s&%s#%s", tmpUrl[0], strings.TrimLeft(param, "?#"), tmpUrl[1])
	}
	if questionMarkIndex > alarmMarkIndex { //?在#号之后,如：https://h5.xxx.com/#/packages/pages/index?code=2&_from=xxx&_goto_time=162
		return fmt.Sprintf("%s&%s", url, strings.TrimLeft(param, "?#"))
	}
	return fmt.Sprintf("%s?%s", url, strings.TrimLeft(param, "?#"))
}

// 是否空串 或者 全是 空格
func IsBlank(str string) bool {
	strLen := len(str)
	if str == "" || strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(str[i])) == false {
			return false
		}
	}
	return true
}
