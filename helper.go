package gosupport

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// 输出，类似php的echo()函数
func Echo(args ...interface{}) {
	_, _ = fmt.Print(args...)
}

func EchoF(args interface{}) {
	_, _ = fmt.Printf("%#v" + GO_EOL, args)
}

func Exit(status int) {
	os.Exit(status)
}

func Die(status int) {
	os.Exit(status)
}

//获取环境变量值
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
//返回红色，用于在终端打印红色字，调用示例：fmt.Println(gosupport.ToRed("hello world"))
func ToRed(str string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", str)
}

func ToGreen(str string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", str)
}

func ToYellow(str string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", str)
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


/*
n 获取随机字符个数
way 选择参与随机字符串的方式
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
//示例：fmt.Println(gosupport.Int64InSlice(9, []int64{5, 6, 9}))
func Int64InSlice(a int64, list []int64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IntInSlice(a int, list []int) bool {
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



//获取当前go版本
func GetGoVersion() string  {
	return strings.Trim(runtime.Version(), "go")
}

//调用示例：gosupport.GenerateUserAgent("user-service", "1.0.0")
func GenerateUserAgent(appname string, ext ...string) string  {
	appversion := ""
	extString := ""
	if len(ext)>1 {
		appversion = "/" + ext[0]
		newExtData := ext[1:]
		for _,v := range newExtData {
			extString += " " + v
		}

	} else if len(ext) == 1 {
		appversion = "/" + ext[0]
	}

	userAgent := fmt.Sprintf(
		"cjs Golang/%s (%s; %s) %s%s%s", GetGoVersion(), runtime.GOOS, runtime.GOARCH, appname, appversion,extString)
	return userAgent
}

//判断code值是否为0，c := '0'认为是int32类型对应ascii值为48
func IsZeroCode(code interface{}) bool {
	var ret bool = false
	switch v:=code.(type) {
	case float64, float32, int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64:
		//byte是uint8，rune是int32
		if fmt.Sprintf("%v", v) == "0" {
			ret = true
		}

	case string:
		if v == "0" {
			ret = true
		}
	default:
		ret =false
	}

	return ret
}
//判断是否为空或零
func IsEmpty(val interface{}) bool  {
	var ret bool
	switch v:=val.(type) {
	case int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64:
		//byte是uint8，rune是int32
		if fmt.Sprintf("%v", v) == "0" {
			ret = true
		}
	case float64, float32:
		//go中0与0.0,0.00,0.0000均相等
		if v==0 {
			ret = true
		}
	case string:
		if v == "0" || v=="" {
			ret = true
		}
	case bool:
		ret = v == false
	default:
		ref := reflect.ValueOf(v)
		ret = ref.IsZero()
	}

	return ret
}

//换行符转html换行标签
func Nl2br(str string) string {
	return strings.Replace(str, "\n", "<br/>", -1)
}

//获取调用我的函数名，即获取当前方法名，返回 包名.方法名、包名.结构体名.方法名
func GetCallFuncName() string {
	//func Caller(skip int) (pc uintptr, file string, line int, ok bool)
	if pc, _, _, ok := runtime.Caller(1);ok {
		f := runtime.FuncForPC(pc)
		return f.Name()
	}
	return ""
}

func GetCallInfo() map[string]string {
	ret := make(map[string]string )
	if pc, fullName, line, ok := runtime.Caller(1);ok {
		f := runtime.FuncForPC(pc)
		ret["func"] = f.Name()
		ret["fullName"] = fullName
		ret["filename"] =  path.Base(fullName)
		ret["line"] = fmt.Sprint(line)
	}
	return ret
}

//获取网关env配置名，即配置key
func GetGwHostEnvName(gwName string) string  {
	gwName = strings.ToUpper(strings.TrimSpace(gwName))
	domainEnvName := fmt.Sprintf("DOMAIN_%s_GW_SERVICE_HOST", gwName)
	return domainEnvName
}

//===============类型转换方法

//float64类型转int64，丢弃小数部分
func Float64Toint64(fNum float64) (int64, error)  {
	s := fmt.Sprintf("%1.3f", fNum)
	sSlice := strings.SplitN(s, ".", 2)
	ret, err := strconv.ParseInt(sSlice[0], 10, 64)
	return ret, err
}


//float64类型转int，丢弃小数部分
func Float64Toint(fNum float64) (int, error)  {
	s := fmt.Sprintf("%1.3f", fNum)
	sSlice := strings.SplitN(s, ".", 2)
	if ret ,err := strconv.Atoi(sSlice[0]);err == nil {
		return ret,nil
	} else {
		return 0, err
	}
}

func MergeStringMap(list ...map[string]string) map[string]string {
	newObj := map[string]string{}
	for _, m := range list {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}
