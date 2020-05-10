package dbutils

import (
	"fmt"
	"regexp"
	"strconv"
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




