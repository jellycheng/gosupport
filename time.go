package gosupport

import (
	"fmt"
	"time"
)

//该文件函数均以Time开头

const  TIME_FORMAT = "2006-01-02 15:04:05"

//返回当前时间结构体指针类型
func TimeNowPtr() *time.Time {
	t := time.Now()
	return &t
}

/**
 * 调用示例：gosupport.TimeFormat2Date(time.Date(2019, 07, 01, 0, 0, 0, 0, time.UTC))
 * 返回格式为 年/月/日，如：2019/07/01 、 2019/11/28
 */
func TimeFormat2Date(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func TimeFormat2DateWay(t time.Time, way int) string {
	var ret string
	year, month, day := t.Date()
	switch way {
		case 1:
			ret = fmt.Sprintf("%d-%02d-%02d", year, month, day)
		case 2:
			ret = fmt.Sprintf("%d.%02d.%02d", year, month, day)
		case 3:
			ret = fmt.Sprintf("%d年%02d月%02d日", year, month, day)
		case 99:
			fallthrough
		default:
			ret = fmt.Sprintf("%d/%02d/%02d", year, month, day)
	}
	return ret
}

//入参指针类型，返回示例：2019-11-28T01:13:36Z
func Time2TimeStr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.UTC().Format("2006-01-02T15:04:05Z")
}
//入参非指针类型，返回示例：2019-11-28T01:13:36Z
func Time2TimeStr2(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05Z")
}

func TimeStr2Time(t string) time.Time {
	temp, err := time.Parse("2006-01-02", t)
	if err != nil {
		panic(err)
	}
	return temp
}

//入参指针类型，返回格式示例：2019-11-28 09:22:30
func TimePtr2Str(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(TIME_FORMAT)
}

func TimePtr2Str2(t time.Time) string {
	return t.Format(TIME_FORMAT)
}

//当前时间转年月日格式，返回示例：20191128
func TimeNow2String() string {
	t := time.Now()
	timeNow := t.Format("20060102")
	return timeNow
}

//当前时间转日格式，返回示例：28
func TimeNow2Day() int {
	t := time.Now()
	timeNow := t.Format("02")
	return Str2Int(timeNow)
}

//当前时间转月格式，返回示例：11
func TimeNow2Month() int {
	t := time.Now()
	timeNow := t.Format("01")
	return Str2Int(timeNow)
}

//当前时间转年月格式，返回示例：201911
func TimeNow2YearMonth() int {
	t := time.Now()
	timeNow := t.Format("200601")
	return Str2Int(timeNow)
}

//当前时间返回年月日
func TimeNow2YMD() (int, int, int)  {
	year, month, day := time.Now().Date()
	return year, int(month), day
}

//时间戳转日期时间格式,调用示例： gosupport.timestamp2DateTime(1569152644, 7)
func timestamp2DateTime(timestamp int, way int) string  {
	var ret string
	timeObj := time.Unix(int64(timestamp), 0) //将时间戳转为时间格式
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	switch way {
		case 1:
			ret = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
		case 2:
			ret = fmt.Sprintf("%d.%02d.%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
		case 3:
			ret = fmt.Sprintf("%d年%02d月%02d日 %02d时%02d分%02d秒", year, month, day, hour, minute, second)
		case 4:
			ret = fmt.Sprintf("%d.%02d.%02d", year, month, day)
		case 5:
			ret = fmt.Sprintf("%d-%02d-%02d", year, month, day)
		case 6:
			ret = fmt.Sprintf("%d/%02d/%02d", year, month, day)
		case 7:
			ret = fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
		case 99:
				fallthrough
		default:
			ret = fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	}
	return ret
}
