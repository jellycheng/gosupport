package gosupport

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
	"regexp"
)

func Exit(status int) {
	os.Exit(status)
}

func Die(status int) {
	os.Exit(status)
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
		fallthrough
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

//获取默认gopath目录
func DefaultGOPATH() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	if home := os.Getenv(env); home != "" {
		return filepath.Join(home, "go")
	}
	return ""
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

//返回 xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func FromatUUIDString(s string) string {
	pattern01 := `^(.{8})(.{0,4})(.{0,4})(.{0,4})(.{1,})$`
	re, err := regexp.Compile(pattern01)
	if err!=nil {
		return s
	} else {
		newStr := re.ReplaceAllString(s, "${1}-${2}-${3}-${4}-${5}")
		return newStr
	}
}

func Uniq(salt string, isFormat bool) string  {
	ret := fmt.Sprintf("%s:%s:%v", salt, RandStr4Byte(6,1), time.Now().UnixNano())
	ret = Md5V1(ret)
	if isFormat {
		return FromatUUIDString(ret)
	} else {
		return ret
	}

}


