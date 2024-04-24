package gosupport

import "time"

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

func NewMyTime() MyTime {
	c := MyTime{}
	c.loc, c.Error = time.LoadLocation(Shanghai)
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
