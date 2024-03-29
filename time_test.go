package gosupport

import (
	"fmt"
	"testing"
	"time"
)

// go test -run="TestTimeFormat2Date"
func TestTimeFormat2Date(t *testing.T) {
	fmt.Println(TimeFormat2Date(time.Now()))
	fmt.Println(TimeFormat2DateWay(time.Now(), 2))

	fmt.Println(Time2TimeStr(TimeNowPtr()))
	fmt.Println(Time2TimeStr2(time.Now()))

	fmt.Println(TimePtr2Str(TimeNowPtr()))
	fmt.Println(TimePtr2Str2(time.Now()))

	fmt.Println("TimeNow2String: ", TimeNow2String())

	fmt.Println("TimeNow2YearMonth:", TimeNow2YearMonth())
	fmt.Println("TimeNow2Month:", TimeNow2Month())
	fmt.Println("TimeNow2Day:", TimeNow2Day())
	y, m, d := TimeNow2YMD()
	fmt.Println("TimeNow2YMD: ", y, m, d)

	//星期
	fmt.Println(time.Now().Weekday().String()) //Thursday
	fmt.Println(time.Now().YearDay())          //返回一年中的第几天

	//当前时间戳 1574907260
	fmt.Println(time.Now().Unix())

	fmt.Println(TimeToSlice(time.Now()))

	fmt.Println("Timestamp2DateTime:", Timestamp2DateTime(1569152644, 1))
	fmt.Println("TimeNow2Format:", TimeNow2Format("2006.01.02 15:04:05"))

	t1, _ := Strtotime(time.RFC3339, "2021-12-21T18:16:23+08:00")
	fmt.Println(t1) //1640081783
}

func TestDateT(t *testing.T) {
	curNow := time.Now()
	fmt.Println(DateT("Y-m-d H:i:s", curNow))

	fmt.Println(DateT("y-m-d h:ii:ss", curNow))

	fmt.Println(DateT("Y/m/d", curNow))
}

func TestSubDays(t *testing.T) {
	//相差天数
	hour, _ := time.ParseDuration("-25h") //25小时前
	t2 := time.Now().In(GetShanghaiTimezone()).Add(hour)
	fmt.Println(SubDays(time.Now().In(GetShanghaiTimezone()), t2))

}

func TestTodayStartEndTime(t *testing.T) {
	s, e := TodayStartEndTime(GetShanghaiTimezone())
	fmt.Println(s.Format(TimeFormat), e.Format(TimeFormat))
	fmt.Println(s.Unix(), e.Unix())

	dayTime, _ := time.ParseInLocation("2006-01-02", "2019-11-09", GetShanghaiTimezone())
	s2, e2 := DayStartEndTime(dayTime)
	fmt.Println(s2.Format(TimeFormat), e2.Format(TimeFormat))
	fmt.Println(s2.Unix(), e2.Unix())

}

// go test -run="TestNewAssertTime"
func TestNewAssertTime(t *testing.T) {
	tmp := NewAssertTime()
	tmp.TimeFormats = append(tmp.TimeFormats, "06年1月2日", "2006年1月2日 15时4分5秒")
	tObj, err := tmp.ParseAssertFormat("2021年7月27日 19时22分09秒", GetShanghaiTimezone())
	if err == nil {
		fmt.Println(tObj.Format(TimeFormat))
	} else {
		fmt.Println(err)
	}

}

// go test -run="TestTimestamp2Time"
func TestTimestamp2Time(t *testing.T) {
	t2 := Timestamp2Time(1640081783)
	fmt.Println(t2.Format("2006-01-02 15:04:05"))
}
