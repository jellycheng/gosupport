package gosupport

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFormat2Date(t *testing.T) {
	fmt.Println(TimeFormat2Date(time.Now()))
	fmt.Println(TimeFormat2DateWay(time.Now(),2))

	fmt.Println(Time2TimeStr(TimeNowPtr()))
	fmt.Println(Time2TimeStr2(time.Now()))

	fmt.Println(TimePtr2Str(TimeNowPtr()))
	fmt.Println(TimePtr2Str2(time.Now()))

	fmt.Println(TimeNow2String())

	fmt.Println(TimeNow2YearMonth())
	fmt.Println(TimeNow2Month())
	fmt.Println(TimeNow2Day())
	fmt.Println(TimeNow2YMD())

	//星期
	fmt.Println(time.Now().Weekday().String()) //Thursday
	fmt.Println(time.Now().YearDay()) //返回一年中的第几天

	//当前时间戳 1574907260
	fmt.Println(time.Now().Unix())

	fmt.Println(timestamp2DateTime(1569152644, 7))

}

func TestDateT(t *testing.T) {
	curNow := time.Now()
	fmt.Println(DateT("Y-m-d H:i:s", curNow))

	fmt.Println(DateT("y-m-d h:ii:ss", curNow))

	fmt.Println(DateT("Y/m/d", curNow))
}

func TestSubDays(t *testing.T) {
	//相差天数
	hour,_ := time.ParseDuration("-25h") //25小时前
	t2 := time.Now().Add(hour)
	fmt.Println(SubDays(time.Now(), t2))

}
