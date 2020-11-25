package dbutils

import (
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
