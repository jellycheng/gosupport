package gosupport

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

/*
  url中追加参数
  调用示例：gosupport.UrlDeal("nfangbian.com/fangan/index/?xyz=1#ab", "a=1&b=2")
 */
func UrlDeal(reqUrl string, otherGetParam string) string {
	if otherGetParam == "" {
		return reqUrl
	}
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


/*
  对特殊字符使用中划线替换
  CreateAnchor("abc？你好?中=国123abc") 返回 abc-你好-中-国123abc
  CreateAnchor("你好中国123abc")返回 你好中国123abc
  CreateAnchor("你好中国 123 abc") 返回 你好中国-123-abc
  CreateAnchor("how 你好 中国123a!bc#de") 返回 how-你好-中国123a-bc-de
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

//转为小驼峰格式： 空格的首字母转大写，如：hello world_abc厉害 转 helloWorld_abc厉害
func ToCamelCase(str string) string {
	str = strings.TrimSpace(str)
	if utf8.RuneCountInString(str) < 2 {
		return str
	}
	var buff strings.Builder
	var temp string
	for _, r := range str {
		c := string(r)
		if c != " " {
			if temp == " " {
				c = strings.ToUpper(c)
			}
			_, _ = buff.WriteString(c)
		}
		temp = c
	}
	return buff.String()
}


//转为snake格式: 全部转小写，空格转_，如：Abc_Xy z_eLsW中国 转 abc_xy_z_elsw中国
func ToSnakeCase(str string) string {
	str = strings.TrimSpace(strings.ToLower(str))
	return strings.Replace(str, " ", "_", -1)
}


//下划线写法转为驼峰写法，如：img_key 转 ImgKey
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
//驼峰转下化线
func Camel2Case(name string) string {
	var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)
	sb := wordBarrierRegex.ReplaceAll(
		[]byte(name),
		[]byte(`${1}_${2}`),
	)
	return string(sb)
}

//首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

//首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}


//获取指定charset字符集下指定长度length的随机字符串
func GetRandStringWithCharset(length int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//获取指定长度的随机字符串
func GetRandString(length int) string {
	return GetRandStringWithCharset(length, CharsetStr1)
}

// 去掉空格后，截取指定长度的字数，并打上...点字符
func GetSummary(s string, length int) string {
	summary := strings.TrimSpace(s)
	if utf8.RuneCountInString(summary) > length {
		summary = Substr(summary, 0, length)
		summary += "..."
	}
	return summary
}

func ToBool(s string) bool {
	s = strings.ToLower(s)
	return s == "true" || s == "yes" || s == "on" || s == "1"
}

func ToBoolOr(s string, defaultValue bool) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return defaultValue
	}
	return b
}

//忽略大小写比较字符串
func EqualsIgnoreCase(a, b string) bool {
	return a == b || strings.ToUpper(a) == strings.ToUpper(b)
}
