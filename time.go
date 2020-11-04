package gosupport

import (
	"fmt"
	"strings"
	"time"
)

//该文件函数大部分以Time开头

const  TIME_FORMAT = "2006-01-02 15:04:05"

func Time() int64 {
	return time.Now().Unix()
}

func Strtotime(format, strtime string) (int64, error) {
	t, err := time.Parse(format, strtime)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}

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


// Y年，4位
// y年，后2位
// m月份，有前导0
// n月份，无加前导0
// d日，有前导0
// j日，无加前导0
// H 24小时制，有前导0
// G 24小时制，无前导0
// h 12小时制，有前导0
// g 12小时制，无前导0
// i 分钟，有前导0
// ii 分钟，无前导0
// s 秒，有前导0
// ss 秒，无前导0
//类似php的写法 Y-m-d H:i:s
func DateT(format string, t time.Time) string {
	res := strings.Replace(format, "Y", t.Format("2006"), -1)
	res = strings.Replace(res, "y", t.Format("06"), -1)
	res = strings.Replace(res, "m", t.Format("01"), -1)
	res = strings.Replace(res, "n", t.Format("1"), -1)
	res = strings.Replace(res, "d", t.Format("02"), -1)
	res = strings.Replace(res, "j", t.Format("2"), -1)

	res = strings.Replace(res, "H", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "G", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "h", t.Format("03"), -1)
	res = strings.Replace(res, "g", t.Format("3"), -1)
	res = strings.Replace(res, "ii", t.Format("4"), -1)
	res = strings.Replace(res, "i", t.Format("04"), -1)
	res = strings.Replace(res, "ss", t.Format("5"), -1)
	res = strings.Replace(res, "s", t.Format("05"), -1)

	return res
}



/**
1分钟以内显示为：刚刚
1小时以内显示为：N分钟前
当天以内显示为：今天 N点N分（如：今天 22:33）
昨天时间显示为：昨天 N点N分（如：昨天 10:15）
在今年显示为：N月N日 N点N分（如：02月03日 09:33）
今年以前显示为：N年N月N日 N点N分（如：2020年09月18日 15:59）
*/
func SubTimeStr(t2 time.Time) string {
	var ret string
	t1 := time.Now()
	t1UnixTime := t1.Unix() //当前时间
	t2UnixTime := t2.Unix()
	if t1UnixTime<t2UnixTime {
		ret = t2.Format("2006-01-02 15:04:05")
		return ret
	}
	subVal := t1UnixTime - t2UnixTime
	if subVal <= 60 {
		ret = "刚刚"
	} else if subVal <= 60 * 60 {
		ret = fmt.Sprintf("%d分钟前", subVal/60)
	} else if t1.Format("20060102") == t2.Format("20060102") {
		ret = fmt.Sprintf("今天 %s", t2.Format("15:04"))
	} else if t1.Format("20060102") == time.Unix(t2.Unix() + 86400, 0).Format("20060102") {
		ret = fmt.Sprintf("昨天 %s", t2.Format("15:04"))
	} else if t1.Format("2006") == t2.Format("2006") {
		ret = t2.Format("01-02 15:04")
	} else {
		ret = t2.Format("2006-01-02 15:04:05")
	}
	return ret
}


/**
 * 已运行时长: d天h小时m分钟s秒
 */
func AlreadyTimeStr(t2 time.Time) string {
	var ret string
	t1 := time.Now()
	t1UnixTime := t1.Unix() //当前时间秒
	t2UnixTime := t2.Unix()
	if t1UnixTime<t2UnixTime {
		ret = "时间倒挂了"
		return ret
	}
	var day int64 //天
	var hour int64 //小时
	var minute int64 //分
	var second int64 //秒
	//总时差
	subVal := t1UnixTime - t2UnixTime
	if subVal/86400>0 { //超过1天
		day = subVal/86400
	}
	if subVal%86400 >0 { //1天内
		hour = (subVal%86400)/3600 //小时
	}
	if subVal%3600>0 { //1小时内
		minute = (subVal%3600)/60 //分
	}
	if subVal % 60>0 {
		second = subVal % 60
	}
	ret = fmt.Sprintf("%d天%d小时%d分%d秒", day,hour,minute,second)
	return ret
}

// 计算日期相差多少天：t1-t2
func SubDays(t1, t2 time.Time) (day int) {
	day = int(t1.Sub(t2).Hours() / 24)
	return
}
