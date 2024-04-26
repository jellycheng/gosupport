package gosupport

import (
	"time"
)

type MyTime struct {
	time  time.Time
	loc   *time.Location
	Error error
}

func (c MyTime) IsZero() bool {
	return c.time.IsZero()
}

func (c MyTime) StdTime() time.Time {
	return c.time.In(c.loc)
}

func (c MyTime) Format(layout string) string {
	return c.StdTime().Format(layout)
}

// 有效时间
func (c MyTime) IsValid() bool {
	if c.Error != nil {
		return false
	}
	if c.time.IsZero() {
		return false
	}
	// 大于零值时间
	if c.StdTime().Unix() > -62135596800 {
		return true
	}
	return false
}

// 无效时间
func (c MyTime) IsInvalid() bool {
	return !c.IsValid()
}

// 时间戳,单位秒
func (c MyTime) Timestamp() int64 {
	if c.Error != nil {
		return 0
	}
	return c.StdTime().Unix()
}

func (c MyTime) Location() string {
	if c.Error != nil {
		return ""
	}
	return c.loc.String()
}

// 获取时区
func (c MyTime) Timezone() string {
	if c.Error != nil {
		return ""
	}
	n, _ := c.StdTime().Zone()
	return n
}

// N年前
func (c MyTime) SubYears(years int) MyTime {
	if c.IsInvalid() {
		return c
	}
	return c.AddYears(-years)
}

// N年后
func (c MyTime) AddYears(years int) MyTime {
	if c.IsInvalid() {
		return c
	}
	c.time = c.StdTime().AddDate(years, 0, 0)
	return c
}

// N个月前
func (c MyTime) SubMonths(months int) MyTime {
	return c.AddMonths(-months)
}

// N个月后
func (c MyTime) AddMonths(months int) MyTime {
	if c.IsInvalid() {
		return c
	}
	c.time = c.StdTime().AddDate(0, months, 0)
	return c
}

// 1天前
func (c MyTime) SubDay() MyTime {
	return c.SubDays(1)
}

// N天前
func (c MyTime) SubDays(days int) MyTime {
	return c.AddDays(-days)
}

// 1天后
func (c MyTime) AddDay() MyTime {
	return c.AddDays(1)
}

// N天后
func (c MyTime) AddDays(days int) MyTime {
	if c.IsInvalid() {
		return c
	}
	c.time = c.StdTime().AddDate(0, 0, days)
	return c
}

// N小时后
func (c MyTime) AddHours(hours int) MyTime {
	if c.IsInvalid() {
		return c
	}
	td := time.Duration(hours) * time.Hour
	c.time = c.StdTime().Add(td)
	return c
}

// N小时前
func (c MyTime) SubHours(hours int) MyTime {
	return c.AddHours(-hours)
}

// N分钟后
func (c MyTime) AddMinutes(minutes int) MyTime {
	if c.IsInvalid() {
		return c
	}
	td := time.Duration(minutes) * time.Minute
	c.time = c.StdTime().Add(td)
	return c
}

// N分钟前
func (c MyTime) SubMinutes(minutes int) MyTime {
	return c.AddMinutes(-minutes)
}

// N秒钟后
func (c MyTime) AddSeconds(seconds int) MyTime {
	if c.IsInvalid() {
		return c
	}
	td := time.Duration(seconds) * time.Second
	c.time = c.StdTime().Add(td)
	return c
}

// N秒钟前
func (c MyTime) SubSeconds(seconds int) MyTime {
	return c.AddSeconds(-seconds)
}

// 获取当前季度
func (c MyTime) Quarter() (quarter int) {
	if c.Error != nil {
		return
	}
	month := c.StdTime().Month()
	switch {
	case month >= 10:
		quarter = 4
	case month >= 7:
		quarter = 3
	case month >= 4:
		quarter = 2
	case month >= 1:
		quarter = 1
	}
	return
}

func (c MyTime) Now(timezone ...string) MyTime {
	if len(timezone) > 0 {
		c.loc, c.Error = time.LoadLocation(timezone[0])
	}
	if c.Error != nil {
		return c
	}
	c.time = time.Now().In(c.loc)
	return c
}

// 明天此刻
func (c MyTime) Tomorrow(timezone ...string) MyTime {
	if len(timezone) > 0 {
		c.loc, c.Error = time.LoadLocation(timezone[0])
	}
	if c.Error != nil {
		return c
	}
	if !c.IsZero() {
		return c.AddDay()
	}
	return c.Now().AddDay()
}

// 昨天
func (c MyTime) Yesterday(timezone ...string) MyTime {
	if len(timezone) > 0 {
		c.loc, c.Error = time.LoadLocation(timezone[0])
	}
	if c.Error != nil {
		return c
	}
	if !c.IsZero() {
		return c.SubDay()
	}
	return c.Now().SubDay()
}

// 周一日期
func (c MyTime) ToMondayString() string {
	ret := ""
	offset := int(time.Monday - c.StdTime().Weekday())
	if offset > 0 {
		offset = -6
	}
	ret = c.time.AddDate(0, 0, offset).Format(DateFormat)
	return ret
}

// 周日日期
func (c MyTime) ToSundayString() string {
	ret := ""
	offset := int(7 - c.StdTime().Weekday())
	if offset == 7 {
		offset = 0
	}
	ret = c.time.AddDate(0, 0, offset).Format(DateFormat)
	return ret
}

func NewMyTime(timezone ...string) MyTime {
	c := MyTime{}
	if len(timezone) > 0 {
		c.loc, c.Error = time.LoadLocation(timezone[0])
	} else {
		c.loc, c.Error = time.LoadLocation(Shanghai)
	}
	return c
}

// 今天此刻
func Now(timezone ...string) MyTime {
	return NewMyTime().Now(timezone...)
}

// 明天此刻
func Tomorrow(timezone ...string) MyTime {
	return NewMyTime().Tomorrow(timezone...)
}

// 昨天此刻
func Yesterday(timezone ...string) MyTime {
	return NewMyTime().Yesterday(timezone...)
}

// 上周周一日期
func PrevWeekMonday() string {
	d := Now().ToMondayString()
	t1, _ := time.Parse(DateFormat, d)
	return t1.AddDate(0, 0, -7).Format(DateFormat)
}

// 上周周日日期
func PrevWeekSunday() string {
	d := Now().ToMondayString()
	t1, _ := time.Parse(DateFormat, d)
	return t1.AddDate(0, 0, -1).Format(DateFormat)
}
