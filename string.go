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
	if index := strings.Index(reqUrl, "#"); index >= 0 {
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

func TrimSpace(str string) string {
	return strings.TrimSpace(str)
}

func TrimSpace4StringSlice(data []string) []string {
	ret := make([]string, 0, len(data))
	for _, item := range data {
		ret = append(ret, strings.TrimSpace(item))
	}
	return ret
}

func CommaSplitString(str string) []string {
	ret := []string{}
	if str == "" {
		return ret
	}
	str = strings.ReplaceAll(str, "，", ",")
	ret = strings.Split(str, ",")
	return ret
}

// 下划线写法转为驼峰写法，如：img_key 转 ImgKey
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 驼峰转下划线，如：ImgKeyXyz 转 Img_Key_Xyz
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

// 首字母转大写
func Ucfirst4PHP(s string) string {
	if len(s) == 0 {
		return ""
	}
	data := []byte(s)
	if data[0] >= 'a' && data[0] <= 'z' {
		data[0] -= 32
	}

	return string(data)
}

// 首字母转小写
func Lcfirst4PHP(s string) string {
	if len(s) == 0 {
		return ""
	}
	data := []byte(s)
	if data[0] >= 'A' && data[0] <= 'Z' {
		data[0] += 32
	}
	return string(data)
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

// 隐藏手机号中间4位
func ReplaceStar4Phone(phone string) string {
	return ReplaceStar4String(phone, 3, 4)
}

// 从开始位置替换为*号 start开始位置-从0开始计算，maxLen最多打星个数
func ReplaceStar4String(str string, start int, maxLen int) string {
	b := []rune(str)
	l := len(b)
	if l == 0 || maxLen == 0 {
		return str
	}
	if start == 0 {
		if maxLen >= l {
			return strings.Repeat("*", l)
		} else {
			return strings.Repeat("*", maxLen) + string(b[maxLen:])
		}
	}
	if start >= l { // 开始位置超出长度
		return str
	}
	var realMax = maxLen
	if l-start < maxLen {
		realMax = l - start
	}
	return string(b[:start]) + strings.Repeat("*", realMax) + string(b[realMax+start:])
}

// column列从1开始
func GetExcelLetter(column int) string {
	result := ""
	if column < 1 {
		return result
	}
	var Letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
		"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z"}
	column = column - 1
	result = Letters[column%26]
	column = column / 26
	for column > 0 {
		column = column - 1
		result = Letters[column%26] + result
		column = column / 26
	}
	return result
}

// 反转字符串
func Strrev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 获取字符串长度，类似php的strlen()函数
func Strlen(str string) int {
	return len(str)
}

// 获取字符串长度，类似php的mb_strlen()函数
func MbStrlen(str string) int {
	return utf8.RuneCountInString(str)
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

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func IsString(v interface{}) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}

func GetBigLetter() []string {
	ret := make([]string, 0)
	for i := 0; i < 26; i++ {
		ret = append(ret, string(rune(65+i)))
	}
	return ret
}

func GetSmallLetter() []string {
	ret := make([]string, 0)
	for i := 0; i < 26; i++ {
		ret = append(ret, string(rune(97+i)))
	}
	return ret
}

// GetExcelNo 下标从0开始，A-Z，AA-AZ，AAA-AAZ
func GetExcelNo(column int) string {
	ret := ""
	start := 65
	i := column % 26
	column = column / 26
	if column >= 1 {
		ret += GetExcelNo(column - 1)
	}
	return ret + string(rune(start+i))
}
