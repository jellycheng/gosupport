package gosupport

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//标识性接口
type IdentifyInterface interface {

}

/*
 * 获取环境变量值
 */
func GetEnv(env, defaultValue string) string {
	v := os.Getenv(env)
	if v == "" {
		return defaultValue
	}
	return v
}

//求和
func IntSum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

//空函数，什么也不做
func Void()  {
	
}

//转成json字符串
func ToJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func MyAssert(guard bool, str string) {
	if !guard {
		panic(str)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Str2Int(str string) int {
	if i, err := strconv.Atoi(str); err != nil {
		return 0
	} else {
		return i
	}
}


/**
n获取随机字符个数
way选择参与随机字符串的方式
*/
func RandStr4Byte(n int, way int) string {
	ret := ""
	if(n<=0) {
		return ret
	}
	var letterStr []byte
	switch way {
	case 2:
		letterStr = []byte("abcdefghijklmnopqrstuvwxyz")
	case 3:
		letterStr = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	case 4:
		letterStr = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	case 5:
		letterStr = []byte("0123456789")
	case 6:
		letterStr = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	case 7:
		letterStr = []byte("abcdefghjkmnpqrstwxyz23456789")
	case 1:
	default:
		letterStr = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	}
	lenStr := len(letterStr)
	b := make([]byte, n)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = letterStr[rand.Intn(lenStr)]
	}
	ret = string(b)
	return ret
}

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)
	if start < 0 || start > length {
		return ""
	}
	if end < 0 {
		return ""
	}
	if end > length {
		end = length
	}
	return string(rs[start:end])
}

//获取变量or值的类型
func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

