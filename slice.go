package gosupport

import (
	"fmt"
)

//切片相关的处理方法

//字符串切片内容去重
func RemoveRepeatByString(data []string) []string  {
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


//int切片内容去重
func RemoveRepeatByInt(data []int) []int  {
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

//int8切片内容去重
func RemoveRepeatByInt8(data []int8) []int8  {
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

//int16切片内容去重
func RemoveRepeatByInt16(data []int16) []int16  {
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

//int32切片内容去重
func RemoveRepeatByInt32(data []int32) []int32  {
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

//int64切片内容去重
func RemoveRepeatByInt64(data []int64) []int64  {
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

//float32切片内容去重
func RemoveRepeatByFloat32(data []float32) []float32  {
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

//float64切片内容去重
func RemoveRepeatByFloat64(data []float64) []float64  {
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

//切片错误结构体
type SliceError struct {
	Msg string
}

func (e *SliceError) Error() string {
	return e.Msg
}

//切片错误内容格式化
func SliceErrorf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return &SliceError{msg}
}

/**
	使用示例
	abc,_ := RemoveRepeatContent([]float64{9.9, 18.9, 22, 9.9, 22, 33})
	if v,ok := abc.([]float64);ok{
		fmt.Println(v[2]) //22
	}
 */
func RemoveRepeatContent(data interface{}) (interface{},error)  {
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