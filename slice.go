package gosupport

import (
	"fmt"
	"strconv"
	"strings"
)

//切片相关的处理方法

// 字符串切片内容去重
func RemoveRepeatByString(data []string) []string {
	ret := make([]string, 0, len(data))
	temp := map[string]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			ret = append(ret, item)
		}
	}
	return ret
}

// int切片内容去重
func RemoveRepeatByInt(data []int) []int {
	ret := make([]int, 0, len(data))
	temp := map[int]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			ret = append(ret, item)
		}
	}
	return ret
}

// int8切片内容去重
func RemoveRepeatByInt8(data []int8) []int8 {
	ret := make([]int8, 0, len(data))
	temp := map[int8]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			ret = append(ret, item)
		}
	}
	return ret
}

// int16切片内容去重
func RemoveRepeatByInt16(data []int16) []int16 {
	ret := make([]int16, 0, len(data))
	temp := map[int16]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			ret = append(ret, item)
		}
	}
	return ret
}

// int32切片内容去重
func RemoveRepeatByInt32(data []int32) []int32 {
	ret := make([]int32, 0, len(data))
	temp := map[int32]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			ret = append(ret, item)
		}
	}
	return ret
}

// int64切片内容去重
func RemoveRepeatByInt64(data []int64) []int64 {
	ret := make([]int64, 0, len(data))
	temp := map[int64]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			ret = append(ret, item)
		}
	}
	return ret
}

// float32切片内容去重
func RemoveRepeatByFloat32(data []float32) []float32 {
	ret := make([]float32, 0, len(data))
	temp := map[float32]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			ret = append(ret, item)
		}
	}
	return ret
}

// float64切片内容去重
func RemoveRepeatByFloat64(data []float64) []float64 {
	ret := make([]float64, 0, len(data))
	temp := map[float64]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			ret = append(ret, item)
		}
	}
	return ret
}

// 切片错误结构体
type SliceError struct {
	Msg string
}

func (e *SliceError) Error() string {
	return e.Msg
}

// 切片错误内容格式化
func SliceErrorf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return &SliceError{msg}
}

/*
使用示例
abc,_ := RemoveRepeatContent([]float64{9.9, 18.9, 22, 9.9, 22, 33})

	if v,ok := abc.([]float64);ok{
		fmt.Println(v[2]) //22
	}
*/
func RemoveRepeatContent(data interface{}) (interface{}, error) {
	switch slice := data.(type) {
	case []string:
		ret := RemoveRepeatByString(slice)
		return ret, nil
	case []int:
		ret := RemoveRepeatByInt(slice)
		return ret, nil
	case []int8:
		ret := RemoveRepeatByInt8(slice)
		return ret, nil
	case []int16:
		ret := RemoveRepeatByInt16(slice)
		return ret, nil
	case []int32:
		ret := RemoveRepeatByInt32(slice)
		return ret, nil
	case []int64:
		ret := RemoveRepeatByInt64(slice)
		return ret, nil
	case []float32:
		ret := RemoveRepeatByFloat32(slice)
		return ret, nil
	case []float64:
		ret := RemoveRepeatByFloat64(slice)
		return ret, nil
	default:
		err := SliceErrorf("Unknown type: %T", slice)
		return nil, err
	}

}

// 把切片内容拼接成字符串
func SliceJointoString(data interface{}, sep string, isRemove bool) (string, error) {
	var tmpData interface{}
	if isRemove {
		if res, err := RemoveRepeatContent(data); err != nil {
			return "", err
		} else {
			tmpData = res
		}
	} else {
		tmpData = data
	}
	var tmpStringSlice []string
	switch slice := tmpData.(type) {
	case []string:
		for _, v := range slice {
			tmpStringSlice = append(tmpStringSlice, fmt.Sprint(v))
		}
	case []int:
		for _, v := range slice {
			tmpStringSlice = append(tmpStringSlice, fmt.Sprint(v))
		}
	case []int8:
		for _, v := range slice {
			tmpStringSlice = append(tmpStringSlice, fmt.Sprint(v))
		}
	case []int16:
		for _, v := range slice {
			tmpStringSlice = append(tmpStringSlice, fmt.Sprint(v))
		}
	case []int32:
		for _, v := range slice {
			tmpStringSlice = append(tmpStringSlice, fmt.Sprint(v))
		}
	case []int64:
		for _, v := range slice {
			tmpStringSlice = append(tmpStringSlice, fmt.Sprint(v))
		}
	case []float32:
		for _, v := range slice {
			tmpStringSlice = append(tmpStringSlice, fmt.Sprint(v))
		}
	case []float64:
		for _, v := range slice {
			tmpStringSlice = append(tmpStringSlice, fmt.Sprint(v))
		}
	default:
		err := SliceErrorf("Unknown type: %T", slice)
		return "", err
	}
	return strings.Join(tmpStringSlice, sep), nil
}

// s := gosupport.InterfacetoString([]interface{}{10, 200}, ",")
func InterfacetoString(data interface{}, sep string) string {
	var tmpStringSlice []string
	for _, v := range data.([]interface{}) {
		tmpStringSlice = append(tmpStringSlice, fmt.Sprint(v))
	}
	str := strings.Join(tmpStringSlice, sep)
	return str
}

// 将[]string转为[]int, s,_ := gosupport.StringSliceToIntSlice([]string{"123", "78919"})
func StringSliceToIntSlice(arr []string) ([]int, bool) {
	result := make([]int, 0)
	for _, i := range arr {
		res, err := strconv.Atoi(i)
		if err != nil {
			return result, false
		}
		result = append(result, res)
	}
	return result, true
}

// 将[]int转为[]string, s,_ := gosupport.IntSliceToStringSlice([]int{123, 78919})
func IntSliceToStringSlice(arr []int) ([]string, bool) {
	result := make([]string, 0)
	for _, i := range arr {
		result = append(result, strconv.Itoa(i))
	}
	return result, true
}

type MySliceInt64 struct {
	data []int64
}

// obj := NewMySliceInt64()  obj.Push(9)
func NewMySliceInt64() *MySliceInt64 {
	obj := MySliceInt64{}
	return &obj
}

// 向栈中添加数据
func (m *MySliceInt64) Push(value int64) {
	m.data = append(m.data, value)
}

// 从栈中获取数据
func (m *MySliceInt64) Pop() (int64, bool) {
	ok := false
	var val int64 = 0
	if len(m.data) > 0 {
		val = m.data[len(m.data)-1]
		m.data = m.data[:len(m.data)-1]
		ok = true
		return val, ok
	} else {
		return val, ok
	}
}

func (m *MySliceInt64) GetData() []int64 {
	return m.data
}

func (m *MySliceInt64) SetData(d []int64) {
	m.data = d
}

func (m *MySliceInt64) RemoveRepeatData() []int64 {
	ret := make([]int64, 0, len(m.data))
	temp := map[int64]struct{}{}
	for _, item := range m.data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			ret = append(ret, item)
		}
	}
	return ret
}

// 过滤字符串切片内容
// 示例：
//
//	 s1 := []string{"hello", "", "yes", "0","yes"}
//		s2 := gosupport.StringFilter(s1, func(s string) bool {
//			if s == "" || s == "0" {
//				return false
//			}
//			return true
//		})
func StringFilter(data []string, callBack func(s string) bool) []string {
	ret := make([]string, 0, len(data))
	for _, item := range data {
		if callBack(item) {
			ret = append(ret, item)
		}
	}
	return ret
}

func IntFilter(data []int, callBack func(i int) bool) []int {
	ret := make([]int, 0, len(data))
	for _, item := range data {
		if callBack(item) {
			ret = append(ret, item)
		}
	}
	return ret
}

func Int64Filter(data []int64, callBack func(i int64) bool) []int64 {
	ret := make([]int64, 0, len(data))
	for _, item := range data {
		if callBack(item) {
			ret = append(ret, item)
		}
	}
	return ret
}

func StringNumber2IntSlice(data []string) []int {
	ret := make([]int, 0, len(data))
	for _, item := range data {
		if i, err := strconv.Atoi(item); err == nil {
			ret = append(ret, i)
		}
	}
	return ret
}

func StringNumber2Int64Slice(data []string) []int64 {
	ret := make([]int64, 0, len(data))
	for _, item := range data {
		if i, err := strconv.ParseInt(item, 10, 64); err == nil {
			ret = append(ret, i)
		}
	}
	return ret
}
