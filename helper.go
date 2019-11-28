package gosupport

import (
	"encoding/json"
	"os"
	"strconv"
)

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


func Str2Int(str string) int {
	if i, err := strconv.Atoi(str); err != nil {
		return 0
	} else {
		return i
	}
}
