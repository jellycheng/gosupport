package gosupport

import (
	"fmt"
	"strconv"
)

//使用示例： gosupport.StrTo("123456789").MustUint64()

type StrTo string

func (f StrTo) String() string {
	return string(f)
}

func (f StrTo) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

func (f StrTo) MustUint8() uint8 {
	v, _ := f.Uint8()
	return v
}

func (f StrTo) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(f.String(), 10, 64)
	return v, err
}

func (f StrTo) MustUint64() uint64 {
	v, _ := f.Uint64()
	return v
}

func (f StrTo) Int() (int, error) {
	if i, err := strconv.Atoi(f.String()); err != nil {
		return 0,err
	} else {
		return i,nil
	}
}

func (f StrTo) IntV0() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 0)
	return int(v), err
}

func (f StrTo) MustInt() int {
	v, _ := f.Int()
	return v
}

func (f StrTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	return int64(v), err
}

func (f StrTo) MustInt64() int64 {
	v, _ := f.Int64()
	return v
}

func (f StrTo) Float32() (float32, error) {
	v, err := strconv.ParseFloat(f.String(), 32)
	return float32(v), err
}

func (f StrTo) MustFloat32() float32 {
	v, _ := f.Float32()
	return v
}

func (f StrTo) Float64() (float64, error) {
	v, err := strconv.ParseFloat(f.String(), 64)
	return v, err
}

func (f StrTo) MustFloat64() float64 {
	v, _ := f.Float64()
	return v
}

//任意值转字符串
func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', ArgInt(args).Get(0, -1), ArgInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', ArgInt(args).Get(0, -1), ArgInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), ArgInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), ArgInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), ArgInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), ArgInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, ArgInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), ArgInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), ArgInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), ArgInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), ArgInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, ArgInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

type ArgInt []int
//示例：var abc gosupport.ArgInt = []int{101, 11,12,33}
//	fmt.Println(abc.Get(1, 10)) //11
//i表取切片第i个值，小于0则取args的第1个值
func (a ArgInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}
