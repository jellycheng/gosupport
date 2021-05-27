package dbutils

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetHashOrd(str string) int64 {
	var ret int64
	isMatch,_ := regexp.MatchString("^[1-9][0-9]*$", str)
	if isMatch {
		//string到int64
		val, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			ret = val
		}
	} else {
		for _,val := range str {//rune方式
			//fmt.Printf("%d,%c \n", val, val)
			strOrd := fmt.Sprintf("%d", val)
			ord, err := strconv.ParseInt(strOrd, 10, 64)
			if err == nil {
				ret += ord
			}
		}
	}

	return ret
}

//返回0～127
func GetHashOrd127(str string) int64 {
	ret := GetHashOrd(str)
	return ret%128
}

func WrapField(field string) string  {
	return "`" + strings.Trim(field, "`") + "`"
}

func WrapTable(field string) string  {
	return "`" + strings.Trim(field, "`") + "`"
}

func WenHaoPlaceholders(n int) string {
	var buf strings.Builder
	for i := 0; i < n-1; i++ {
		buf.WriteString("?,")
	}
	if n > 0 {
		buf.WriteString("?")
	}
	return buf.String()
}

// db字段类型转go类型, fieldTypeStr=bigint(20)
func FiledType2GoType(fieldTypeStr string)  string{
	fieldTypeStr = strings.TrimSpace(fieldTypeStr)
	typeArr := strings.Split(fieldTypeStr,"(")
	realType := typeArr[0]
	if isMatch,_ := regexp.MatchString(`\s+`, typeArr[0]);isMatch {
		typeArr2 := strings.Split(typeArr[0]," ")
		realType = typeArr2[0]
	}
	switch realType {
	case "int":
		return "int64"
	case "integer":
		return "int"
	case "mediumint":
		return "int"
	case "bit":
		return "int"
	case "year":
		return "int"
	case "smallint":
		return "int"
	case "tinyint":
		return "int"
	case "bigint":
		return "int64"
	case "decimal":
		return "float64"
	case "double":
		return "float64"
	case "float":
		return "float32"
	case "real":
		return "float32"
	case "numeric":
		return "float32"
	case "timestamp":
		return "time.Time"
	case "datetime":
		return "time.Time"
	case "time":
		return "time.Time"
	default:
		return "string"
	}
}

// md5加密
func Md5V1(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	has := h.Sum([]byte(""))
	md5str := fmt.Sprintf("%x", has)
	return md5str
}


func PinInStr(strs []string) string {
	var buf strings.Builder
	for k, str := range strs {
		if k == 0 {
			buf.WriteString("'" + str + "'")
		} else {
			buf.WriteString(",'" + str + "'")
		}
	}
	inStr := buf.String()
	return inStr
}


func JoinInt642Str(data []int64, sep string) string {
	switch len(data) {
	case 0:
		return ""
	case 1:
		return strconv.FormatInt(data[0], 10)
	}
	var buf strings.Builder
	for k, v := range data {
		s := strconv.FormatInt(v, 10)
		if k == 0 {
			buf.WriteString(s)
		} else {
			buf.WriteString(sep + s)
		}
	}
	inStr := buf.String()
	return inStr
}

func JoinInt2Str(data []int, sep string) string {
	switch len(data) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(data[0])
	}
	var buf strings.Builder
	for k, v := range data {
		s := strconv.Itoa(v)
		if k == 0 {
			buf.WriteString(s)
		} else {
			buf.WriteString(sep + s)
		}
	}
	inStr := buf.String()
	return inStr
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// _my_field_name_2 转成 XMyFieldName_2.
func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		t = append(t, 'X')
		i++
	}

	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		if isASCIILower(c) {
			c ^= ' '
		}
		t = append(t, c)
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

//list切片中是否有a字符串
func StrInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetLikeStr(str string, pos int) string {
	ret := ""
	switch pos {
	case 0:
		ret = fmt.Sprintf("%%%s%%", str)
	case 1:
		ret = fmt.Sprintf("%%%s", str)
	case 2:
		ret = fmt.Sprintf("%s%%", str)
	default:
		ret = fmt.Sprintf("%%%s%%", str)
	}

	return ret
}
