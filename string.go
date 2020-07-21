package gosupport

import (
	"strings"
	"unicode"
)

/**
 * url中追加参数
 * gosupport.UrlDeal("nfangbian.com/fangan/index/?xyz=1#ab", "a=1&b=2")
 */
func UrlDeal(reqUrl string, otherGetParam string) string {
	ret := ""
	suffix := ""
	if index := strings.Index(reqUrl, "#");index>=0 {
		ret = reqUrl[:index]
		suffix = reqUrl[index:]
	} else {
		ret = reqUrl
	}
	if strings.ContainsRune(reqUrl, '?') {
		ret += "&" + otherGetParam + suffix
	} else {
		ret += "?" + otherGetParam + suffix
	}
	return ret
}


/**
 * 对特殊字符使用中划线替换
 * CreateAnchor("abc？你好?中=国123abc") 返回 abc-你好-中-国123abc
 * CreateAnchor("你好中国123abc")返回 你好中国123abc
 * CreateAnchor("你好中国 123 abc") 返回 你好中国-123-abc
 * CreateAnchor("how 你好 中国123a!bc#de") 返回 how-你好-中国123a-bc-de
 */
func CreateAnchor(str string) string {
	var anchorName []rune
	var futureDash = false
	for _, r := range str {
		switch {
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			if futureDash && len(anchorName) > 0 {
				anchorName = append(anchorName, '-')
			}
			futureDash = false
			anchorName = append(anchorName, unicode.ToLower(r))
		default:
			futureDash = true
		}
	}
	return string(anchorName)
}

