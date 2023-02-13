package gosupport

import (
	"errors"
	"fmt"
	"regexp"
)

// IsAccountName 判断是否普通账号名，必须以字母开头，可由字母、数字、下划线组成
func IsAccountName(str string) bool {
	isMatch, err := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_]*$", str)
	if err != nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}

// IsMail 是否邮箱
func IsMail(mail string) bool {
	isMatch, err := regexp.MatchString("^([.a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((.[a-zA-Z0-9_-]{2,10}){1,3})$", mail)
	if err != nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}

// IsMobile 是否手机号，11位数字
func IsMobile(str string) bool {
	isMatch, err := regexp.MatchString("^1[0-9]{10}$", str)
	if err != nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}

// IsPhone 是否座机
func IsPhone(str string) bool {
	return checkRegexp(str, "^([0-9]{3,4}-)?[0-9]{7,8}$")
}

// IsUrl 是否链接
func IsUrl(str string) bool {
	return checkRegexp(str, `^http[s]?://.*`)
}

func checkRegexp(val string, reg string) bool {
	isMatch, err := regexp.MatchString(reg, val)
	if err != nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}

// RegexpVerify 正则表达式验证字符串
func RegexpVerify(val string, reg string) bool {
	return checkRegexp(val, reg)
}

// IsNumber 字符串是否为正整数数字字符串
func IsNumber(str string) bool {
	isMatch, err := regexp.MatchString("^[1-9][0-9]*$", str)
	if err != nil {
		return false
	}
	if isMatch {
		return true
	} else if str == "0" {
		return true
	} else {
		return false
	}
}

// IsFloatNumber 字符串是否为浮点数字符串
func IsFloatNumber(str string) bool {
	isMatch, err := regexp.MatchString("^[0-9]+[.]?[0-9]*$", str)
	if err != nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}

// StripTags 去除html标签
func StripTags(s string, tags ...string) string {
	if len(tags) == 0 {
		tags = append(tags, "")
	}
	for _, tag := range tags {
		stripTagsRe := regexp.MustCompile(`(?i)<\/?` + tag + `[^<>]*>`)
		s = stripTagsRe.ReplaceAllString(s, "")
	}
	return s
}

// ExtractContent4Tag 从str中提取tag内的内容
func ExtractContent4Tag(str, tag string) []string {
	ret := make([]string, 0)
	tagRe := regexp.MustCompile(`(?ims)<` + tag + `.*?[^<>]*>(.*?)</\s*` + tag + `\s*>`)
	resTmp := tagRe.FindAllStringSubmatch(str, -1)
	for _, v := range resTmp {
		ret = append(ret, v[1])
	}
	return ret
}

// GetWwVerifyVal 提取内容，例如 WW_verify_P3fNz9uLSkgAlsnI.txt 提取出P3fNz9uLSkgAlsnI
func GetWwVerifyVal(str string) string {
	ret := ""
	rObj := regexp.MustCompile(WwVerify)
	tmp := rObj.FindStringSubmatch(str)
	if len(tmp) == 2 {
		ret = tmp[1]
	}
	return ret
}

// GetMpVerifyVal 提取内容，例如 MP_verify_d4RP2dwJOG3lDBub.txt 提取出 d4RP2dwJOG3lDBub
func GetMpVerifyVal(str string) string {
	ret := ""
	rObj := regexp.MustCompile(MpVerify)
	tmp := rObj.FindStringSubmatch(str)
	if len(tmp) == 2 {
		ret = tmp[1]
	}
	return ret
}

// extractCode 提取{code}中的code值
func extractCode(src string) []string {
	ret := []string{}
	pattern01 := `{(\w+)}`
	re := regexp.MustCompile(pattern01)
	tmp := re.FindAllStringSubmatch(src, -1)
	for _, v := range tmp {
		ret = append(ret, v[1])
	}
	return ret
}

// Replace4code 把src串中{code}替换为val值
func Replace4code(src, code, val string) (string, error) {
	ret := ""
	pattern01 := fmt.Sprintf(`{%s}`, code)
	re, err := regexp.Compile(pattern01)
	if err != nil {
		return ret, errors.New("正则表达式错误: " + pattern01)
	} else {
		ret = re.ReplaceAllString(src, val)
		return ret, nil
	}

}

func Replace4Map(src string, data map[string]string) string {
	ret := src
	// 提取code
	codes := extractCode(src)
	if len(codes) == 0 {
		return ret
	}
	for _, c := range codes {
		if val, ok := data[c]; ok {
			ret, _ = Replace4code(ret, c, val)
		}
	}
	return ret
}
