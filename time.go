package gosupport

import (
	"fmt"
	"time"
)

//该文件函数均以Time开头

const  TIME_FORMAT = "2006-01-02 15:04:05"

/**
 * gosupport.TimeFormat2Date(time.Date(2019, 07, 01, 0, 0, 0, 0, time.UTC))
 * 返回格式为 年/月/日，如：2019/07/01
 */
func TimeFormat2Date(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}


func Time2TimeStr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.UTC().Format("2006-01-02T15:04:05Z")
}

func TimeStr2Time(t string) time.Time {
	temp, err := time.Parse("2006-01-02", t)
	if err != nil {
		panic(err)
	}
	return temp
}

func TimePtr2Str(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(TIME_FORMAT)
}

func TimeNowPtr() *time.Time {
	t := time.Now()
	return &t
}

func TimeNow2String() string {
	t := time.Now()
	timeNow := t.Format("20060102")
	return timeNow
}

func TimeNow2Month() int {
	t := time.Now()
	timeNow := t.Format("200601")
	return Str2Int(timeNow)
}

